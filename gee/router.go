package gee

import "net/http"

// router 定义一个路由,存在所有匹配模式和处理方法
type router struct {
	handlers map[string]HandlerFunc
}

// newRouter 创建时路由实例
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

// addRoute 添加（注册）路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle 分发路由处理
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, c.Path)
	}
}
