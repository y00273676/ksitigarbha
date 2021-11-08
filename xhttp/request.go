package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var hostKeyUpperCase = strings.ToUpper("Host")

type request struct {
	Client    *http.Client
	Request   *http.Request
	Transport *http.Transport
	Err       error
}

type response struct {
	Response *http.Response
	Body     []byte
	Err      error
}

type Request interface {
	Get(url string) Request
	PostJSON(url string, data interface{}) Request
	PostForm(url string, data url.Values) Request
	SetBody(body io.ReadCloser) Request
	HttpRequest() *http.Request
	Head(url string) Request
	Options(url string) Request
	Put(url string) Request
	Delete(url string) Request
	AddHeader(key, value string) Request
	AddHeaders(headers map[string]string) Request
	AddCookie(key, value string) Request
	AddCookiesByCookies(cookies []*http.Cookie) Request
	AddCookiesByMaps(cookies []map[string]string) Request
	AddProxy(proxy string) Request
	AddUserAgent(value string) Request
	AddContentType(value string) Request
	AddReferer(value string) Request
	AddHost(value string) Request
	AddTimeout(value time.Duration) Request
	AddTLSHandshakeTimeout(value time.Duration) Request
	AddResponseHeaderTimeout(value time.Duration) Request
	AddExpectContinueTimeout(value time.Duration) Request
	AddIdleConnTimeout(value time.Duration) Request
	SetRoundTripper(roundTripper http.RoundTripper) Request
	CleanCookie() Request
	BasicAuth(username, password string) Request
	DoWithCtx(ctx context.Context) (Response, error)
	Do() (Response, error)
}

type Response interface {
	Cookies() []*http.Cookie
	Headers() map[string]string
	URLToString() string
	ReadBody() Response
	ToBytes() []byte
	ToString() string
	ToMap() map[string]interface{}
	HttpResponse() *http.Response
	Close() error
	Status() string
	StatusCode() int
	Error() error
	IsTimeout() bool
}

func NewRequest() Request {
	return &request{
		Client: &http.Client{},
		Request: &http.Request{
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		},
		Transport: DefaultTransport,
	}
}

// DefaultTransport 是http.DefaultTransport， 不指定的情况下http用的是这个实现
var DefaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   2 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func (r *request) parseUrl(rawurl string) *url.URL {
	urlObj, err := url.Parse(rawurl)
	if err != nil {
		r.Err = err
	}
	return urlObj
}

func (r *request) Get(url string) Request {
	r.Request.Method = "GET"
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) PostJSON(url string, data interface{}) Request {
	r.Request.Method = "POST"
	r.AddHeader("Content-Type", "application/json")
	if data != nil {
		jsonByte, err := json.Marshal(data)
		if err != nil {
			r.Err = err
		}
		reader := bytes.NewReader(jsonByte)
		r.Request.Body = ioutil.NopCloser(reader)
	}
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) PostForm(url string, data url.Values) Request {
	r.Request.Method = "POST"
	r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	r.Request.URL = r.parseUrl(url)
	r.Request.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))
	return r
}

func (r *request) SetBody(body io.ReadCloser) Request {
	r.Request.Body = body
	return r
}

func (r *request) HttpRequest() *http.Request {
	return r.Request
}

func (r *request) Head(url string) Request {
	r.Request.Method = "HEAD"
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) Options(url string) Request {
	r.Request.Method = "OPTIONS"
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) Put(url string) Request {
	r.Request.Method = "PUT"
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) Delete(url string) Request {
	r.Request.Method = "DELETE"
	r.Request.URL = r.parseUrl(url)
	return r
}

func (r *request) AddHeader(key, value string) Request {
	if strings.ToUpper(key) == hostKeyUpperCase {
		r.AddHost(value)
	} else {
		r.Request.Header.Set(key, value)
	}
	return r
}

func (r *request) AddHeaders(headers map[string]string) Request {
	for key, value := range headers {
		r.AddHeader(key, value)
	}
	return r
}

func (r *request) AddCookie(key, value string) Request {
	r.Request.AddCookie(&http.Cookie{Name: key, Value: value})
	return r
}

func (r *request) AddCookiesByCookies(cookies []*http.Cookie) Request {
	for _, cookie := range cookies {
		r.Request.AddCookie(cookie)
	}
	return r
}

func (r *request) AddCookiesByMaps(cookies []map[string]string) Request {
	for _, cookie := range cookies {
		for key, value := range cookie {
			r.AddCookie(key, value)
		}
	}
	return r
}

func (r *request) AddProxy(proxy string) Request {
	r.Transport.Proxy = func(i *http.Request) (*url.URL, error) { return url.Parse(proxy) }
	return r
}

func (r *request) AddUserAgent(value string) Request {
	r.Request.Header.Set("User-Agent", value)
	return r
}

func (r *request) AddContentType(value string) Request {
	r.Request.Header.Set("Content-Type", value)
	return r
}

func (r *request) AddReferer(value string) Request {
	r.Request.Header.Set("Referer", value)
	return r
}

// Golang 中设置 Host 这个 header.不能用直接添加 Header 的方式不会生效. 需要直接修改 req.Host 参数
// 相关 issue. https://github.com/golang/go/issues/29865
func (r *request) AddHost(value string) Request {
	r.Request.Host = value
	return r
}

func (r *request) AddTimeout(value time.Duration) Request {
	r.Client.Timeout = value
	return r
}

func (r *request) AddTLSHandshakeTimeout(value time.Duration) Request {
	r.Transport.TLSHandshakeTimeout = value
	return r
}

func (r *request) AddResponseHeaderTimeout(value time.Duration) Request {
	r.Transport.ResponseHeaderTimeout = value
	return r
}

func (r *request) AddExpectContinueTimeout(value time.Duration) Request {
	r.Transport.ExpectContinueTimeout = value
	return r
}

func (r *request) AddIdleConnTimeout(value time.Duration) Request {
	r.Transport.IdleConnTimeout = value
	return r
}

func (r *request) CleanCookie() Request {
	r.Request.Header.Set("Cookie", "")
	return r
}

func (r *request) BasicAuth(username, password string) Request {
	r.Request.SetBasicAuth(username, password)
	return r
}
func (r *request) SetRoundTripper(roundTripper http.RoundTripper) Request {
	r.Client.Transport = roundTripper
	return r
}

// chooseTransport 能够用户自定义一个Transport. 在有自定义的Transport的时候
// 也就是Client里已经有了一个Transport则用用户手动指定的，否则用外部指定的
func (r *request) chooseTransport() {
	if r.Client.Transport == nil {
		r.Client.Transport = r.Transport
	}
}

func (r *request) DoWithCtx(ctx context.Context) (Response, error) {
	if r.Err != nil {
		return &response{&http.Response{}, nil, r.Err}, r.Err
	}
	r.Request = r.Request.WithContext(ctx)
	r.chooseTransport()
	resp, err := r.Client.Do(r.Request)
	if err != nil {
		r.Err = err
	}
	return &response{resp, nil, err}, err
}

func (r *request) Do() (Response, error) {
	if r.Err != nil {
		return &response{&http.Response{}, nil, r.Err}, r.Err
	}
	r.chooseTransport()
	resp, err := r.Client.Do(r.Request)
	if err != nil {
		r.Err = err
	}
	return &response{resp, nil, err}, err
}

func (r *response) Cookies() []*http.Cookie {
	return r.Response.Cookies()
}

func (r *response) Headers() map[string]string {
	headers := make(map[string]string)
	for key, value := range r.Response.Header {
		headers[key] = strings.Join(value, ";")
	}
	return headers
}

func (r *response) URLToString() string {
	return r.Response.Request.URL.String()
}

func (r *response) ReadBody() Response {
	if r.Err != nil {
		return r
	}
	resByte, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		r.Err = err
	}
	r.Body = resByte
	r.Close()
	return r
}

func (r *response) ToBytes() []byte {
	if r.Err != nil {
		return nil
	}

	if len(r.Body) == 0 {
		r.ReadBody()
	}
	return r.Body
}

func (r *response) ToString() string {
	if r.Err != nil {
		return ""
	}

	if len(r.Body) == 0 {
		r.ReadBody()
	}
	return string(r.Body)
}

func (r *response) ToMap() map[string]interface{} {
	if r.Err != nil {
		return nil
	}

	if len(r.Body) == 0 {
		r.ReadBody()
	}
	var result map[string]interface{}
	if r.Err != nil {
		return result
	}
	err := json.Unmarshal(r.Body, &result)
	if err != nil {
		r.Err = err
		return nil
	}
	return result
}

func (r *response) HttpResponse() *http.Response {
	return r.Response
}

func (r *response) Close() error {

	if r.Response == nil {
		return nil
	}
	if r.Response.Body != nil {
		return r.Response.Body.Close()
	}
	return nil
}

func (r *response) Status() string {
	return r.Response.Status
}

func (r *response) StatusCode() int {
	return r.Response.StatusCode
}

func (r *response) Error() error {
	return r.Err
}

func (r *response) IsTimeout() bool {
	e, isTimeout := r.Err.(interface {
		Timeout() bool
	})
	if isTimeout {
		return e.Timeout()
	}
	return false
}
