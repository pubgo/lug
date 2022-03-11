// Code generated by protoc-gen-resty. DO NOT EDIT.
// versions:
// - protoc-gen-resty v0.1.0
// - protoc           v3.19.4
// source: proto/gid/echo_service.proto

package gid

import (
	context "context"
	v2 "github.com/go-resty/resty/v2"
	go_json "github.com/goccy/go-json"
	reflect "reflect"
)

type EchoServiceResty interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error)
	// EchoPatch method receives a NonStandardUpdateRequest and returns it.
	EchoPatch(ctx context.Context, in *DynamicMessageUpdate, opts ...func(req *v2.Request)) (*DynamicMessageUpdate, error)
	// EchoUnauthorized method receives a simple message and returns it. It must
	// always return a google.rpc.Code of `UNAUTHENTICATED` and a HTTP Status code
	// of 401.
	EchoUnauthorized(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error)
}

func NewEchoServiceResty(client *v2.Client) EchoServiceResty {
	client.SetContentLength(true)
	return &echoServiceResty{client: client}
}

type echoServiceResty struct {
	client *v2.Client
}

func (c *echoServiceResty) Echo(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error) {
	var req = c.client.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	for i := range opts {
		opts[i](req)
	}
	var body map[string]interface{}
	if in != nil {
		var rv = reflect.ValueOf(in).Elem()
		var rt = reflect.TypeOf(in).Elem()
		for i := 0; i < rt.NumField(); i++ {
			if val, ok := rt.Field(i).Tag.Lookup("param"); ok && val != "" {
				req.SetPathParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("query"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
				continue
			}
			if body == nil {
				body = make(map[string]interface{})
			}
			if val, ok := rt.Field(i).Tag.Lookup("json"); ok && val != "" {
				body[val] = rv.Field(i).String()
			}
		}
	}
	req.SetBody(body)
	var resp, err = req.Execute("POST", "/v1/example/echo/{id}")
	if err != nil {
		return nil, err
	}
	out := new(SimpleMessage)
	if err := go_json.Unmarshal(resp.Body(), out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceResty) EchoBody(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error) {
	var req = c.client.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	for i := range opts {
		opts[i](req)
	}
	var body map[string]interface{}
	if in != nil {
		var rv = reflect.ValueOf(in).Elem()
		var rt = reflect.TypeOf(in).Elem()
		for i := 0; i < rt.NumField(); i++ {
			if val, ok := rt.Field(i).Tag.Lookup("param"); ok && val != "" {
				req.SetPathParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("query"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
				continue
			}
			if body == nil {
				body = make(map[string]interface{})
			}
			if val, ok := rt.Field(i).Tag.Lookup("json"); ok && val != "" {
				body[val] = rv.Field(i).String()
			}
		}
	}
	req.SetBody(body)
	var resp, err = req.Execute("POST", "/v1/example/echo_body")
	if err != nil {
		return nil, err
	}
	out := new(SimpleMessage)
	if err := go_json.Unmarshal(resp.Body(), out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceResty) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error) {
	var req = c.client.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	for i := range opts {
		opts[i](req)
	}
	var body map[string]interface{}
	if in != nil {
		var rv = reflect.ValueOf(in).Elem()
		var rt = reflect.TypeOf(in).Elem()
		for i := 0; i < rt.NumField(); i++ {
			if val, ok := rt.Field(i).Tag.Lookup("param"); ok && val != "" {
				req.SetPathParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("query"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
				continue
			}
			if body == nil {
				body = make(map[string]interface{})
			}
			if val, ok := rt.Field(i).Tag.Lookup("json"); ok && val != "" {
				body[val] = rv.Field(i).String()
			}
		}
	}
	req.SetBody(body)
	var resp, err = req.Execute("DELETE", "/v1/example/echo_delete")
	if err != nil {
		return nil, err
	}
	out := new(SimpleMessage)
	if err := go_json.Unmarshal(resp.Body(), out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceResty) EchoPatch(ctx context.Context, in *DynamicMessageUpdate, opts ...func(req *v2.Request)) (*DynamicMessageUpdate, error) {
	var req = c.client.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	for i := range opts {
		opts[i](req)
	}
	var body map[string]interface{}
	if in != nil {
		var rv = reflect.ValueOf(in).Elem()
		var rt = reflect.TypeOf(in).Elem()
		for i := 0; i < rt.NumField(); i++ {
			if val, ok := rt.Field(i).Tag.Lookup("param"); ok && val != "" {
				req.SetPathParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("query"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
				continue
			}
			if body == nil {
				body = make(map[string]interface{})
			}
			if val, ok := rt.Field(i).Tag.Lookup("json"); ok && val != "" {
				body[val] = rv.Field(i).String()
			}
		}
	}
	req.SetBody(body)
	var resp, err = req.Execute("PATCH", "/v1/example/echo_patch")
	if err != nil {
		return nil, err
	}
	out := new(DynamicMessageUpdate)
	if err := go_json.Unmarshal(resp.Body(), out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceResty) EchoUnauthorized(ctx context.Context, in *SimpleMessage, opts ...func(req *v2.Request)) (*SimpleMessage, error) {
	var req = c.client.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	for i := range opts {
		opts[i](req)
	}
	if in != nil {
		var rv = reflect.ValueOf(in).Elem()
		var rt = reflect.TypeOf(in).Elem()
		for i := 0; i < rt.NumField(); i++ {
			if val, ok := rt.Field(i).Tag.Lookup("param"); ok && val != "" {
				req.SetPathParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("query"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
				continue
			}
			if val, ok := rt.Field(i).Tag.Lookup("json"); ok && val != "" {
				req.SetQueryParam(val, rv.Field(i).String())
			}
		}
	}
	var resp, err = req.Execute("GET", "/v1/example/echo_unauthorized")
	if err != nil {
		return nil, err
	}
	out := new(SimpleMessage)
	if err := go_json.Unmarshal(resp.Body(), out); err != nil {
		return nil, err
	}
	return out, nil
}
