package HttpPipeline

import (
	"php-in-go/Include/Http"
)

type HttpPipeline struct {
	request   *Http.Request
	response  *Http.Response
	pipelines []func(*Http.Request, *Http.Response, func(*Http.Request, *Http.Response))
}

func NewHttpPipeline() *HttpPipeline {
	return &HttpPipeline{}
}

func (h *HttpPipeline) Send(request *Http.Request, response *Http.Response) *HttpPipeline {
	h.request = request
	h.response = response
	return h
}

func (h *HttpPipeline) Through(unit ...func(*Http.Request, *Http.Response, func(*Http.Request, *Http.Response))) *HttpPipeline {
	h.pipelines = unit
	return h
}

func (h *HttpPipeline) reverse() []func(*Http.Request, *Http.Response, func(*Http.Request, *Http.Response)) {
	for from, to := 0, len(h.pipelines)-1; from < to; from, to = from+1, to-1 {
		h.pipelines[from], h.pipelines[to] = h.pipelines[to], h.pipelines[from]
	}
	return h.pipelines
}

func (h *HttpPipeline) Then(p func(*Http.Request, *Http.Response)) {
	f := h.reduce(
		h.pipelines,
		func(
			f func(*Http.Request, *Http.Response),
			f2 func(
				*Http.Request,
				*Http.Response,
				func(*Http.Request, *Http.Response),
			),
		) func(*Http.Request, *Http.Response) {
			return func(request *Http.Request, response *Http.Response) {
				f2(request, response, f)
			}
		},
		p,
	)
	f(h.request, h.response)
}

func (h *HttpPipeline) reduce(
	s []func(
		*Http.Request,
		*Http.Response,
		func(
			*Http.Request,
			*Http.Response,
		),
	),
	f func(
		func(*Http.Request, *Http.Response),
		func(
			*Http.Request,
			*Http.Response,
			func(*Http.Request, *Http.Response),
		),
	) func(*Http.Request, *Http.Response),
	init func(*Http.Request, *Http.Response),
) func(*Http.Request, *Http.Response) {
	acc := init
	for _, t := range s {
		acc = f(acc, t)
	}
	return acc
}
