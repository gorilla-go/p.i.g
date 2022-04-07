package Http

import "net/http"

type Response struct {
	responseWriter http.ResponseWriter
}

func BuildResponse(responseWriter http.ResponseWriter) *Response {
	return &Response{responseWriter: responseWriter}
}

func (r *Response) SetCode(code int) *Response {
	r.responseWriter.WriteHeader(code)
	return r
}

func (r *Response) Write(content string) *Response {
	_, err := r.responseWriter.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	return r
}
