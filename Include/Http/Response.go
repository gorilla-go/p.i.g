package Http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"php-in-go/Include/Foundation/Util"
)

type Response struct {
	responseWriter http.ResponseWriter
}

func BuildResponse(responseWriter http.ResponseWriter) *Response {
	return &Response{responseWriter: responseWriter}
}

func (r *Response) setCode(code int) *Response {
	r.responseWriter.WriteHeader(code)
	return r
}

func (r *Response) SetHeader(key string, content string) *Response {
	r.responseWriter.Header().Set(key, content)
	return r
}

func (r *Response) write(content string) {
	_, err := r.responseWriter.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

func (r *Response) Html(content string) {
	r.SetHeader("content-type", "text/html; charset=utf-8")
	r.setCode(http.StatusOK)
	r.write(content)
}

func (r *Response) HtmlWithCode(content string, code int) {
	r.SetHeader("content-type", "text/html; charset=utf-8")
	r.setCode(code)
	r.write(content)
}

func (r *Response) Json(j interface{}) {
	r.SetHeader("content-type", "application/json; charset=utf-8")
	r.setCode(http.StatusOK)
	marshal, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	r.write(string(marshal))
}

func (r *Response) AddCookie(cookie *http.Cookie) {
	r.SetHeader("set-cookie", cookie.String())
}

func (r *Response) Dump(i interface{}) {
	r.SetHeader("content-type", "text/html; charset=utf-8")
	r.setCode(http.StatusOK)
	r.write(Util.Dump(i, 1))
}

func (r *Response) View(template string, params map[string]interface{}) {

}

func (r *Response) Echo(i interface{}) {
	r.SetHeader("content-type", "text/plain; charset=utf-8")
	r.write(fmt.Sprintf("%v", i))
}

func (r *Response) Redirect(url string, code int) {
	r.SetHeader("Location", url)
	r.setCode(code)
}
