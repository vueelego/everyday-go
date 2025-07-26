package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// 定义上下文Context,用于包装请求上下文
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Method     string
	Path       string
	StatusCode int
}

// newContext 创建是一个上下文
// 在哪里调用创建呢? 那肯定是入口函数 ServeHTTP
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{
		Writer: w,
		Req:    r,

		Method: r.Method,
		Path:   r.URL.Path,
	}

	return ctx
}

// PostForm 解析form并获取对应key的值
func (c *Context) PostForm(key string) string {
	// FormValue 会执行 ParseMultipartForm 解析数据
	return c.Req.FormValue(key)
}

// Query 获取URL的查询字符值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置并写入响应状态
func (c *Context) Status(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

// String 响应字符串类型数据
func (c *Context) String(statusCode int, format string, values ...any) {
	// 1. 设置响应返回的类型
	c.SetHeader("Content-Type", "text/plain")

	// 2. 设置响应状态码
	c.Status(statusCode)

	// 3. 写入数据
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 响应json数据
func (c *Context) JSON(statusCode int, data any) {
	// 1. 设置响应返回的类型
	c.SetHeader("Content-Type", "application/json")

	// 2. 设置响应状态码
	c.Status(statusCode)

	// 3. 写入数据
	err := json.NewEncoder(c.Writer).Encode(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// Data 系统自动识别类型响应数据
func (c *Context) Data(statusCode int, data []byte) {
	c.Status(statusCode)
	c.Writer.Write(data)
}

// HTML 响应html模板
func (c *Context) HTML(statusCode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(statusCode)
	c.Writer.Write([]byte(html))
}
