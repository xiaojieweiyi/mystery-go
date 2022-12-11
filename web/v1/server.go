package web

import "net/http"

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	addRoute(method string, path string, handler HandleFunc)
}

type HTTPServer struct {
	router
}

var _ Server = &HTTPServer{}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resq: writer,
	}
	s.serve(ctx)
}

func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *HTTPServer) Post(path string, handler HandleFunc) {

}

func (s *HTTPServer) Get(path string, handler HandleFunc) {

}

func (s *HTTPServer) serve(ctx *Context) {
	//n, ok := s.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
}
