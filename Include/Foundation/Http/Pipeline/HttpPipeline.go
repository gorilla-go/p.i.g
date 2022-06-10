package HttpPipeline

import (
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
)

type HttpPipeline struct {
	request   *Request.Request
	response  *Response.Response
	pipelines []func(*Request.Request, *Response.Response, func(*Request.Request, *Response.Response))
}

func NewHttpPipeline() *HttpPipeline {
	return &HttpPipeline{}
}

func (h *HttpPipeline) Send(request *Request.Request, response *Response.Response) *HttpPipeline {
	h.request = request
	h.response = response
	return h
}

func (h *HttpPipeline) Through(unit ...func(*Request.Request, *Response.Response, func(*Request.Request, *Response.Response))) *HttpPipeline {
	h.pipelines = unit
	return h
}

func (h *HttpPipeline) reverse() []func(*Request.Request, *Response.Response, func(*Request.Request, *Response.Response)) {
	for from, to := 0, len(h.pipelines)-1; from < to; from, to = from+1, to-1 {
		h.pipelines[from], h.pipelines[to] = h.pipelines[to], h.pipelines[from]
	}
	return h.pipelines
}

func (h *HttpPipeline) Then(p func(*Request.Request, *Response.Response)) {
	f := h.reduce(
		h.pipelines,
		func(
			f func(*Request.Request, *Response.Response),
			f2 func(
				*Request.Request,
				*Response.Response,
				func(*Request.Request, *Response.Response),
			),
		) func(*Request.Request, *Response.Response) {
			return func(request *Request.Request, response *Response.Response) {
				f2(request, response, f)
			}
		},
		p,
	)
	f(h.request, h.response)
}

func (h *HttpPipeline) reduce(
	s []func(
		*Request.Request,
		*Response.Response,
		func(
			*Request.Request,
			*Response.Response,
		),
	),
	f func(
		func(*Request.Request, *Response.Response),
		func(
			*Request.Request,
			*Response.Response,
			func(*Request.Request, *Response.Response),
		),
	) func(*Request.Request, *Response.Response),
	init func(*Request.Request, *Response.Response),
) func(*Request.Request, *Response.Response) {
	acc := init
	for _, t := range s {
		acc = f(acc, t)
	}
	return acc
}
