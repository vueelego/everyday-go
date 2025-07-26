package gee

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
