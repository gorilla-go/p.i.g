package Http

import "net/http"

type Response struct {
	*http.Response
}

func GetNewResponse(response *http.Response) *Response {
	return &Response{
		&http.Response{
			Status:           response.Status,
			StatusCode:       response.StatusCode,
			Proto:            response.Proto,
			ProtoMajor:       response.ProtoMajor,
			ProtoMinor:       response.ProtoMinor,
			Header:           response.Header,
			Body:             response.Body,
			ContentLength:    response.ContentLength,
			TransferEncoding: response.TransferEncoding,
			Close:            response.Close,
			Uncompressed:     response.Uncompressed,
			Trailer:          response.Trailer,
			Request:          response.Request,
			TLS:              response.TLS,
		},
	}
}
