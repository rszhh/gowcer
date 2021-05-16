package module

import "net/http"

// Data 数据的接口类型
type Data interface {
	Valid() bool
}

// 使用Data接口将下面三个类型归为一类

// Request 数据请求的类型
type Request struct {
	// httpReq HTTP请求
	httpReq *http.Request
	// depth 请求的深度
	// 代表一个请求的深度值，用于网络爬虫的自动停止
	depth uint32
}

// NewRequest 创建一个新的请求实例
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{
		httpReq: httpReq,
		depth:   depth,
	}
}

// HTTPReq 获取HTTP请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// Depth 获取当前请求的深度
func (req *Request) Depth() *uint32 {
	return req.Depth()
}

// Valid 用于判断当前请求是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

// Response 代表数据响应的类型
type Response struct {
	// httpResp 代表HTTP响应
	httpResp *http.Response
	// depth 代表响应的深度
	depth uint32
}

// NewResponse 用于创建一个新的响应实例
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

// HTTPResp 用于获取HTTP响应
func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

// Depth 用于获取响应深度
func (resp *Response) Depth() uint32 {
	return resp.depth
}

// Valid 用于判断响应是否有效
func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

// Item 代表条目的类型
type Item map[string]interface{}

// Valid 用于判断条目是否有效
func (item Item) Valid() bool {
	return item != nil
}
