package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

// envelope 方便快速定义响应数据，类似 gin.H
type envelope map[string]interface{}

// writeJSON 快速写入响应
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		slog.Error("解析data错误", "error", err)
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)

	return nil
}

// readJSON 读取请求体的body数据
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 4 << 20 // 4MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	// dec.DisallowUnknownFields() // 调用此方法，如果存在未知字段，Decode 直接返回错误

	// 第一次解码
	err := dec.Decode(dst)
	if err != nil {
		// 简单处理几种常见的错误
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.Is(err, io.EOF):
			return errors.New("请求体不能为空")
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("请求体不过大于 %d", maxBytesError.Limit)
		case errors.As(err, &invalidUnmarshalError):
			// input 传的就不是指针
			panic(err)
		default:
			return err
		}
	}

	// 第二次解码，如果还是可以绑定到一个匿名结构体，说明传了多个JOSN对象
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("请求体不支持多个JSON对象")
	}

	return nil
}

// serverError 服务端错误通用响应
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// errorResponse 错误响应
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(500)
	}
}
