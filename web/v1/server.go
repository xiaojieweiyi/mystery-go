package web

import "net/http"

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	addRouter(method string, path string, handler HandleFunc)
}

type HTTPServer struct {
	router
}

func NewHttpServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
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
	s.addRouter(http.MethodPost, path, handler)
}

func (s *HTTPServer) Get(path string, handler HandleFunc) {
	s.addRouter(http.MethodGet, path, handler)
}

func (s *HTTPServer) serve(ctx *Context) {
	n, ok := s.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || n.handler == nil {
		ctx.Resq.WriteHeader(404)
		ctx.Resq.Write([]byte("Not Found"))
		return
	}
	n.handler(ctx)
}
