package Http

import (
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Params    *url.Values
	AppConfig map[string]interface{}
	*http.Request
}

func BuildRequest(request *http.Request, appConfig map[string]interface{}) *Request {
	return &Request{
		Params:    nil,
		AppConfig: appConfig,
		Request: &http.Request{
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
			Response:         request.Response,
		},
	}
}

func (r *Request) IsAjax() bool {
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func (r *Request) IsPjax() bool {
	if r.IsAjax() == false {
		return false
	}
	_, b := r.ParamVar(r.AppConfig["pjaxName"].(string))
	return b
}

func (r *Request) ParamVar(s string) (string, bool) {
	if r.Params.Has(s) == false {
		return "", false
	}
	return r.Params.Get(s), true
}

func (r *Request) PostVar(s string) interface{} {
	contentType := r.Request.Header.Get("Content-Type")
	if strings.Contains(contentType, "form-data") {
		if err := r.Request.ParseMultipartForm(1024 * 8); err != nil {
			panic(err)
		}
		v := r.Request.MultipartForm.Value
		if val, exist := v[s]; exist && len(val) == 1 {
			return val[0]
		}
		if v, e := v[s+"[]"]; e {
			return v
		}
		return nil
	}

	if strings.Contains(contentType, "x-www-form-urlencoded") {
		if err := r.Request.ParseForm(); err != nil {
			panic(err)
		}
		v := r.Request.PostForm
		if v.Has(s) {
			return v.Get(s)
		}
		if v.Has(s + "[]") {
			return v.Get(s + "[]")
		}
		return nil
	}
	return nil
}
