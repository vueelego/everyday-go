package gee

import (
	"net/http"
	"strings"
)

// router 定义一个路由,存在所有匹配模式和处理方法
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// newRouter 创建时路由实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 解析路由模式，拆分api路径地址；
// 比如 “/api/user/:id” => [api,user,:id]
func parsePattern(pattern string) []string {
	segs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, seg := range segs {
		if seg != "" { // 路由模式经过strings.Split后第一个通常是空字符
			parts = append(parts, seg)
			if seg[0] == '*' { // 比如：*filepath
				break
			}
		}
	}

	return parts
}

// addRoute 添加（注册）路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	// 如果还没有该http method方法节点，则创建
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	// 在指定method节点下添加
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	/* roots 以http的method方法为第一个节点，
		  roots
		/   |   \
	GET  POST PUT  ...
	/|\  /|\  /|\
	.............
	*/
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
