package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pubgo/funk/errors/errutil"
	"github.com/pubgo/funk/log"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	_ "github.com/gofiber/contrib/websocket"
	_ "github.com/gorilla/websocket"
	_ "nhooyr.io/websocket"
)

type streamHTTP struct {
	opts           muxOptions
	ctx            context.Context
	method         *httpPathRule
	w              io.Writer
	wHeader        http.Header
	rbuf           []byte    // stream read buffer
	r              io.Reader //
	rHeader        http.Header
	header         metadata.MD
	trailer        metadata.MD
	params         url.Values
	contentType    string
	accept         string
	acceptEncoding string
	recvCount      int
	sendCount      int
	sentHeader     bool
	hasBody        bool // HTTP method has a body
	rEOF           bool // stream read EOF
}

var _ grpc.ServerStream = (*streamHTTP)(nil)

func (s *streamHTTP) SetHeader(md metadata.MD) error {
	if s.sentHeader {
		return fmt.Errorf("already sent headers")
	}
	s.header = metadata.Join(s.header, md)
	return nil
}

func (s *streamHTTP) SendHeader(md metadata.MD) error {
	if s.sentHeader {
		return fmt.Errorf("already sent headers")
	}
	s.header = metadata.Join(s.header, md)

	h := s.wHeader
	setOutgoingHeader(h, s.header)
	// don't write the header code, wait for the body.
	s.sentHeader = true

	if sh := s.opts.statsHandler; sh != nil {
		sh.HandleRPC(s.ctx, &stats.OutHeader{
			Header:      s.header.Copy(),
			Compression: s.acceptEncoding,
		})
	}
	return nil
}

func (s *streamHTTP) SetTrailer(md metadata.MD) {
	s.trailer = metadata.Join(s.trailer, md)
}

func (s *streamHTTP) Context() context.Context {
	return grpc.NewContextWithServerTransportStream(s.ctx, &serverTransportStream{
		ServerStream: s,
		method:       s.method.grpcMethodName,
	})
}

func (s *streamHTTP) writeMsg(c Codec, b []byte, contentType string) (int, error) {
	count := s.sendCount
	if count == 0 {
		h := s.wHeader
		h.Set("Content-Type", contentType)
		if !s.sentHeader {
			if err := s.SendHeader(nil); err != nil {
				return count, err
			}
		}
	}
	s.sendCount += 1
	if s.method.desc.IsStreamingServer() {
		codec, ok := c.(StreamCodec)
		if !ok {
			return count, fmt.Errorf("codec %s does not support streaming", codec.Name())
		}
		_, err := codec.WriteNext(s.w, b)
		return count, err
	}
	return count, s.opts.writeAll(s.w, b)
}

func (s *streamHTTP) SendMsg(m interface{}) error {
	reply := m.(proto.Message)

	fRsp, ok := s.w.(http.Flusher)
	if ok {
		defer fRsp.Flush()
	}

	cur := reply.ProtoReflect()
	for _, fd := range s.method.rspBody {
		cur = cur.Mutable(fd).Message()
	}
	msg := cur.Interface()

	contentType := s.accept
	c, err := s.getCodec(contentType, cur)
	if err != nil {
		return err
	}

	bytes := bytesPool.Get().(*[]byte)
	b := (*bytes)[:0]
	defer func() {
		if cap(b) < s.opts.maxReceiveMessageSize {
			*bytes = b
			bytesPool.Put(bytes)
		}
	}()

	if cur.Descriptor().FullName() == "google.api.HttpBody" {
		fds := cur.Descriptor().Fields()
		fdContentType := fds.ByName(protoreflect.Name("content_type"))
		fdData := fds.ByName(protoreflect.Name("data"))
		pContentType := cur.Get(fdContentType)
		pData := cur.Get(fdData)

		b = append(b, pData.Bytes()...)
		contentType = pContentType.String()
	} else {
		var err error
		b, err = c.MarshalAppend(b, msg)
		if err != nil {
			return status.Errorf(codes.Internal, "%s: error while marshaling: %v", c.Name(), err)
		}
	}

	if _, err := s.writeMsg(c, b, contentType); err != nil {
		return err
	}

	if stats := s.opts.statsHandler; stats != nil {
		// TODO: raw payload stats.
		stats.HandleRPC(s.ctx, outPayload(false, m, b, b, time.Now()))
	}
	return nil
}

func (s *streamHTTP) readMsg(c Codec, b []byte) (int, []byte, error) {
	if s.rEOF {
		return s.recvCount, nil, io.EOF
	}

	count := s.recvCount
	s.recvCount += 1
	if s.method.desc.IsStreamingClient() {
		codec, ok := c.(StreamCodec)
		if !ok {
			return count, nil, fmt.Errorf("codec %q does not support streaming", codec.Name())
		}
		b = append(b, s.rbuf...)
		b, n, err := codec.ReadNext(b, s.r, s.opts.maxReceiveMessageSize)
		if err == io.EOF {
			s.rEOF, err = true, nil
		}
		s.rbuf = append(s.rbuf[:0], b[n:]...)
		return count, b[:n], err

	}
	b, err := s.opts.readAll(b, s.r)
	if err == io.EOF {
		s.rEOF, err = true, nil
	}
	return count, b, err
}

func (s *streamHTTP) getCodec(mediaType string, cur protoreflect.Message) (Codec, error) {
	codecType := string(cur.Descriptor().FullName())
	if c, ok := s.opts.codecs[codecType]; ok {
		return c, nil
	}
	codecType = mediaType
	if c, ok := s.opts.codecs[codecType]; ok {
		return c, nil
	}
	return nil, status.Errorf(codes.Internal, "no codec registered for content-type %q", mediaType)
}

func (s *streamHTTP) decodeRequestArgs(args proto.Message) (int, error) {
	bytes := bytesPool.Get().(*[]byte)
	b := (*bytes)[:0]
	defer func() {
		if cap(b) < s.opts.maxReceiveMessageSize {
			*bytes = b
			bytesPool.Put(bytes)
		}
	}()

	cur := args.ProtoReflect()
	for _, fd := range s.method.reqBody {
		cur = cur.Mutable(fd).Message()
	}
	msg := cur.Interface()

	c, err := s.getCodec(s.contentType, cur)
	if err != nil {
		return -1, err
	}

	var (
		count int
	)
	count, b, err = s.readMsg(c, b)
	if err != nil {
		return count, err
	}

	if cur.Descriptor().FullName() == "google.api.HttpBody" {
		fds := cur.Descriptor().Fields()
		fdContentType := fds.ByName("content_type")
		fdData := fds.ByName("data")
		cur.Set(fdContentType, protoreflect.ValueOfString(s.contentType))

		cpy := make([]byte, len(b))
		copy(cpy, b)
		cur.Set(fdData, protoreflect.ValueOfBytes(cpy))
	} else {
		if err := c.Unmarshal(b, msg); err != nil {
			return count, status.Errorf(codes.Internal, "%s: error while unmarshaling: %v", c.Name(), err)
		}
	}
	if stats := s.opts.statsHandler; stats != nil {
		// TODO: raw payload stats.
		stats.HandleRPC(s.ctx, inPayload(false, msg, b, b, time.Now()))
	}
	return count, nil
}

func (s *streamHTTP) RecvMsg(m interface{}) error {
	args := m.(proto.Message)

	var count int
	if s.method.hasReqBody && s.hasBody {
		var err error
		count, err = s.decodeRequestArgs(args)
		if err != nil {
			return err
		}
	} else {
		count = s.recvCount
		s.recvCount += 1
		if s.rEOF {
			return io.EOF
		}
		s.rEOF = true
	}
	if count == 0 {
		if err := s.params.set(args); err != nil {
			return err
		}
	}
	return nil
}

func isWebsocketRequest(r *http.Request) bool {
	for _, header := range r.Header["Upgrade"] {
		if header == "websocket" {
			return true
		}
	}
	return false
}

type twirpError struct {
	Code    string            `json:"code"`
	Message string            `json:"msg"`
	Meta    map[string]string `json:"meta"`
}

func (m *Mux) encError(w http.ResponseWriter, r *http.Request, err error) {
	s, _ := status.FromError(err)
	if isTwirp := r.Header.Get("Twirp-Version") != ""; isTwirp {
		accept := "application/json"

		w.Header().Set("Content-Type", accept)
		w.WriteHeader(HTTPStatusCode(s.Code()))

		codeStr := strings.ToLower(code.Code_name[int32(s.Code())])

		terr := &twirpError{
			Code:    codeStr,
			Message: s.Message(),
		}
		b, err := json.Marshal(terr)
		if err != nil {
			panic(err) // ...
		}
		w.Write(b) //nolint
		return
	}

	accept := negotiateContentType(r.Header, m.opts.contentTypeOffers, "application/json")
	c := m.opts.codecs[accept]

	w.Header().Set("Content-Type", accept)
	w.WriteHeader(HTTPStatusCode(s.Code()))

	b, err := c.Marshal(s.Proto())
	if err != nil {
		panic(err) // ...
	}
	w.Write(b) //nolint
}

func (m *Mux) serveHTTP(w http.ResponseWriter, r *http.Request) error {
	ctx, mdata := newIncomingContext(r.Context(), r.Header)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	isWebsocket := isWebsocketRequest(r)

	verb := r.Method
	if isWebsocket {
		verb = kindWebsocket
	}

	methodName := "/" + strings.Trim(strings.TrimSpace(r.URL.Path), "/")
	hds := s.handlers[methodName]
	var (
		hd     *handler
		params params
		err    error
		method = &httpPathRule{grpcMethodName: methodName, hasReqBody: true}
	)

	if len(hds) == 0 {
		method, params, err = s.match(r.URL.Path, verb)
		if err != nil {
			return err
		}

		queryParams, err := method.parseQueryParams(r.URL.Query())
		if err != nil {
			return err
		}
		params = append(params, queryParams...)

		hd, err = s.pickMethodHandler(method.grpcMethodName)
		if err != nil {
			return err
		}
	} else {
		hd = hds[0]
		method.desc = hd.desc
	}

	// Handle stats.
	beginTime := time.Now()

	if isWebsocket {
		var responseHeader http.Header
		if r.Header.Get("Sec-WebSocket-Protocol") != "" {
			responseHeader = http.Header{
				"Sec-WebSocket-Protocol": []string{r.Header.Get("Sec-WebSocket-Protocol")},
			}
		}
		conn, err := upgrade.Upgrade(w, r, responseHeader)
		if err != nil {
			return err
		}

		conn.SetReadLimit(maxMessageSize)
		conn.SetReadDeadline(time.Now().Add(pongWait))
		conn.SetPingHandler(nil)
		conn.SetPongHandler(func(string) error {
			return conn.SetReadDeadline(time.Now().Add(pongWait))
		})
		go func() {
			ticker := time.NewTicker(pingPeriod)
			defer ticker.Stop()
			defer func() {
				log.Info().Msg("close websocket ping")
			}()
			for {
				select {
				case <-ticker.C:
					conn.SetWriteDeadline(time.Now().Add(writeWait))
					if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						log.Err(err).Msg("failed to write ping message")
						break
					}
				case <-ctx.Done():
					return
				}
			}
		}()

		stream := &streamWS{
			ctx:      ctx,
			conn:     conn,
			pathRule: method,
			params:   params,
		}
		herr := hd.handler(&m.opts, stream)
		if herr != nil {
			s, _ := status.FromError(herr)
			// TODO: limit message size.
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(errutil.GrpcCodeToHTTP(s.Code()), s.Message()))
		} else {
			conn.WriteMessage(websocket.CloseMessage, []byte{})
		}

		return nil
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}
	contentEncoding := r.Header.Get("Content-Encoding")

	var body io.Reader = r.Body
	if cz := m.opts.compressors[contentEncoding]; cz != nil {
		z, err := cz.Decompress(r.Body)
		if err != nil {
			return err
		}
		body = z
	}

	accept := negotiateContentType(r.Header, m.opts.contentTypeOffers, contentType)
	acceptEncoding := negotiateContentEncoding(r.Header, m.opts.encodingTypeOffers)

	var resp io.Writer = w
	if cz := m.opts.compressors[acceptEncoding]; cz != nil {
		w.Header().Set("Content-Encoding", acceptEncoding)
		z, err := cz.Compress(w)
		if err != nil {
			return err
		}
		defer z.Close()
		resp = z
	}

	stream := &streamHTTP{
		ctx:    ctx,
		method: method,
		params: params,
		opts:   m.opts,

		// write
		w:       resp,
		wHeader: w.Header(),

		// read
		r:       body,
		rHeader: r.Header,

		contentType:    contentType,
		accept:         accept,
		acceptEncoding: acceptEncoding,
		hasBody:        r.ContentLength > 0 || r.ContentLength == -1,
	}
	herr := hd.handler(&m.opts, stream)
	// Handle stats.
	if herr != nil {
		if !stream.sentHeader {
			w.Header().Set("Content-Encoding", "identity") // try to avoid gzip
		}
		m.encError(w, r, herr)
	}
	return nil
}
