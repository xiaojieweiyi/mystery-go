package web

import "testing"

func TestServer(t *testing.T) {
	s := NewHttpServer()
	s.Get("/", func(ctx *Context) {
		ctx.Resq.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *Context) {
		ctx.Resq.Write([]byte("hello, user"))
	})
	s.Start(":8080")
}
