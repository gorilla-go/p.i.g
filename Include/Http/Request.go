package Http

import "net/http"

type Request struct {
	*http.Request
}

func BuildRequest(request *http.Request) *Request {
	return &Request{
		&http.Request{
			Method:           request.Method,
			URL:              request.URL,
			Proto:            request.Proto,
			ProtoMajor:       request.ProtoMajor,
			ProtoMinor:       request.ProtoMinor,
			Header:           request.Header,
			Body:             request.Body,
			GetBody:          request.GetBody,
			ContentLength:    request.ContentLength,
			TransferEncoding: request.TransferEncoding,
			Close:            request.Close,
			Host:             request.Host,
			Form:             request.Form,
			PostForm:         request.PostForm,
			MultipartForm:    request.MultipartForm,
			Trailer:          request.Trailer,
			RemoteAddr:       request.RemoteAddr,
			RequestURI:       request.RequestURI,
			TLS:              request.TLS,
			Cancel:           request.Cancel,
			Response:         request.Response,
		},
	}
}
