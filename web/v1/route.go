package web

type router struct {
	trees map[string]*node
}

type node struct {
	path     string
	children map[string]*node
	handler  HandleFunc
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

func (r *router) addRouter(method string, path string, handler HandleFunc) {
	if path == "" {
		panic("web: 路由是空字符串")
	}
	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}
	if path != "/" && path[len(path)-1] == '/' {
		path
	}
}
