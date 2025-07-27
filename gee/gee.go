package gee

import (
	"net/http"
)

/* 第一天
0. 定义http处理函数类型
1. 定义一个Engine引擎,包含一个router map对象,存在pattern和对应的处理函数
2. 定义一个addRoute函数,注册路由
3. 分别定义GET/POST 注册路由函数
4. 定义run启动服务
*/

// type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// type Engine struct {
// 	router map[string]HandlerFunc
// }

// func New() *Engine {
// 	engine := &Engine{
// 		router: make(map[string]HandlerFunc),
// 	}
// 	return engine
// }

// func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
// 	e.router[fmt.Sprintf("%s-%s", method, pattern)] = handler
// }

// func (e *Engine) GET(pattern string, handler HandlerFunc) {
// 	e.addRoute(http.MethodGet, pattern, handler)
// }

// func (e *Engine) POST(pattern string, handler HandlerFunc) {
// 	e.addRoute(http.MethodPost, pattern, handler)
// }

// func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	key := r.Method + "-" + r.URL.Path
// 	handler, ok := e.router[key]
// 	if !ok {
// 		http.Error(w, key+": not found", http.StatusNotFound)
// 		return
// 	}
// 	handler(w, r)
// }

// func (e *Engine) Run(addr string) error {
// 	return http.ListenAndServe(addr, e)
// }

// =================================================================================

/** 第二天

 */

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	return engine
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

// ServeHTTP 路由入口执行函数
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 构造 Context 上下文
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
