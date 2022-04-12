package Http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	Util2 "php-in-go/Include/Util"
	"strconv"
)

type Response struct {
	Code           int
	responseWriter http.ResponseWriter
	ErrorStack     string
	ErrorMessage   string
}

func BuildResponse(responseWriter http.ResponseWriter) *Response {
	return &Response{
		responseWriter: responseWriter,
		Code:           200,
	}
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
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
	r.SetCode(http.StatusOK)
	r.write(content)
}

func (r *Response) HtmlWithCode(content string, code int) {
	r.SetHeader("content-type", "text/html; charset=utf-8")
	r.SetCode(code)
	r.write(content)
}

func (r *Response) Json(j interface{}) {
	r.SetHeader("content-type", "application/json; charset=utf-8")
	r.SetCode(http.StatusOK)
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
	r.SetCode(http.StatusOK)
	r.write(Util2.Dump(i, 1))
}

func (r *Response) View(template string, params map[string]interface{}) {

}

func (r *Response) GetResponseWriter() http.ResponseWriter {
	return r.responseWriter
}

func (r *Response) Echo(i interface{}) {
	r.SetHeader("content-type", "text/plain; charset=utf-8")
	r.SetCode(http.StatusOK)
	r.write(fmt.Sprintf("%v", i))
}

func (r *Response) EchoWithCode(i interface{}, code int) {
	r.SetHeader("content-type", "text/plain; charset=utf-8")
	r.SetCode(code)
	r.write(fmt.Sprintf("%v", i))
}

func (r *Response) Download(file string, name string) {
	if Util2.IsFile(file) == false {
		panic("Invalid file path: " + file)
	}
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	r.SetHeader("Content-type", "application/octet-stream")
	r.SetHeader("Accept-Ranges", "bytes")
	r.SetHeader("Accept-Length", strconv.Itoa(len(fileBytes)))
	r.SetHeader("Content-Disposition", "attachment; filename="+name)
	r.SetCode(200)
	_, err = r.responseWriter.Write(fileBytes)
	if err != nil {
		panic(err)
	}
}

func (r *Response) Redirect(url string, code int) {
	r.SetHeader("Location", url)
	r.SetCode(code)
}
