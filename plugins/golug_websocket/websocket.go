package golug_websocket

import (
	"errors"
	"io"
	"sync"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// New returns a new `handler func(*Conn)` that upgrades a client to the
// websocket protocol, you can pass an optional config.
func NewWebsocket(handler func(*Conn), config ...ClientCfg) fiber.Handler {
	// Init config
	var cfg ClientCfg
	if len(config) > 0 {
		cfg = config[0]
	}
	if len(cfg.Origins) == 0 {
		cfg.Origins = []string{"*"}
	}
	if cfg.ReadBufferSize == 0 {
		cfg.ReadBufferSize = 1024
	}
	if cfg.WriteBufferSize == 0 {
		cfg.WriteBufferSize = 1024
	}
	var upGrader = websocket.FastHTTPUpgrader{
		HandshakeTimeout:  cfg.HandshakeTimeout,
		Subprotocols:      cfg.SubProtocols,
		ReadBufferSize:    cfg.ReadBufferSize,
		WriteBufferSize:   cfg.WriteBufferSize,
		EnableCompression: cfg.EnableCompression,
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			if cfg.Origins[0] == "*" {
				return true
			}
			origin := utils.GetString(ctx.Request.Header.Peek("Origin"))
			for i := range cfg.Origins {
				if cfg.Origins[i] == origin {
					return true
				}
			}
			return false
		},
	}
	return func(c *fiber.Ctx) error {
		conn := acquireConn()
		// locals
		c.Context().VisitUserValues(func(key []byte, value interface{}) {
			conn.locals[string(key)] = value
		})

		// params
		params := c.Route().Params
		for i := 0; i < len(params); i++ {
			conn.params[utils.ImmutableString(params[i])] = utils.ImmutableString(c.Params(params[i]))
		}

		// queries
		c.Context().QueryArgs().VisitAll(func(key, value []byte) {
			conn.queries[string(key)] = string(value)
		})

		// cookies
		c.Context().Request.Header.VisitAllCookie(func(key, value []byte) {
			conn.cookies[string(key)] = string(value)
		})

		if err := upGrader.Upgrade(c.Context(), func(fconn *websocket.Conn) {
			conn.Conn = fconn
			defer releaseConn(conn)
			handler(conn)
		}); err != nil { // Upgrading required
			return fiber.ErrUpgradeRequired
		}

		return nil
	}
}

// Conn https://godoc.org/github.com/gorilla/websocket#pkg-index
type Conn struct {
	*websocket.Conn
	locals  map[string]interface{}
	params  map[string]string
	cookies map[string]string
	queries map[string]string
}

// Conn pool
var poolConn = sync.Pool{
	New: func() interface{} {
		return new(Conn)
	},
}

// Acquire Conn from pool
func acquireConn() *Conn {
	conn := poolConn.Get().(*Conn)
	conn.locals = make(map[string]interface{})
	conn.params = make(map[string]string)
	conn.queries = make(map[string]string)
	conn.cookies = make(map[string]string)
	return conn
}

// Return Conn to pool
func releaseConn(conn *Conn) {
	conn.Conn = nil
	poolConn.Put(conn)
}

// Locals makes it possible to pass interface{} values under string keys scoped to the request
// and therefore available to all following routes that match the request.
func (conn *Conn) Locals(key string) interface{} {
	return conn.locals[key]
}

// Params is used to get the route parameters.
// Defaults to empty string "" if the param doesn't exist.
// If a default value is given, it will return that value if the param doesn't exist.
func (conn *Conn) Params(key string, defaultValue ...string) string {
	v, ok := conn.params[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Query returns the query string parameter in the url.
// Defaults to empty string "" if the query doesn't exist.
// If a default value is given, it will return that value if the query doesn't exist.
func (conn *Conn) Query(key string, defaultValue ...string) string {
	v, ok := conn.queries[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Cookies is used for getting a cookie value by key
// Defaults to empty string "" if the cookie doesn't exist.
// If a default value is given, it will return that value if the cookie doesn't exist.
func (conn *Conn) Cookies(key string, defaultValue ...string) string {
	v, ok := conn.cookies[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Constants are taken from https://github.com/fasthttp/websocket/blob/master/conn.go#L43

// Close codes defined in RFC 6455, section 11.7.
const (
	CloseNormalClosure           = 1000
	CloseGoingAway               = 1001
	CloseProtocolError           = 1002
	CloseUnsupportedData         = 1003
	CloseNoStatusReceived        = 1005
	CloseAbnormalClosure         = 1006
	CloseInvalidFramePayloadData = 1007
	ClosePolicyViolation         = 1008
	CloseMessageTooBig           = 1009
	CloseMandatoryExtension      = 1010
	CloseInternalServerErr       = 1011
	CloseServiceRestart          = 1012
	CloseTryAgainLater           = 1013
	CloseTLSHandshake            = 1015
)

// The message types are defined in RFC 6455, section 11.8.
const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

var (
	ErrBadHandshake = errors.New("websocket: bad handshake")
	ErrCloseSent    = errors.New("websocket: close sent")
	ErrReadLimit    = errors.New("websocket: read limit exceeded")
)

// FormatCloseMessage formats closeCode and text as a WebSocket close message.
// An empty message is returned for code CloseNoStatusReceived.
func FormatCloseMessage(closeCode int, text string) []byte {
	return websocket.FormatCloseMessage(closeCode, text)
}

// IsCloseError returns boolean indicating whether the error is a *CloseError
// with one of the specified codes.
func IsCloseError(err error, codes ...int) bool {
	return websocket.IsCloseError(err, codes...)
}

// IsUnexpectedCloseError returns boolean indicating whether the error is a
// *CloseError with a code not in the list of expected codes.
func IsUnexpectedCloseError(err error, expectedCodes ...int) bool {
	return websocket.IsUnexpectedCloseError(err, expectedCodes...)
}

// IsWebSocketUpgrade returns true if the client requested upgrade to the
// WebSocket protocol.
func IsWebSocketUpgrade(c *fiber.Ctx) bool {
	return websocket.FastHTTPIsWebSocketUpgrade(c.Context())
}

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *websocket.Conn, term string) io.Reader {
	return websocket.JoinMessages(c, term)
}
