package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建并注册路由
	mux := &http.ServeMux{}
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	// 启动服务，给予处理者是Logger
	http.ListenAndServe(":8080", &Logger{mux})
}

type Logger struct {
	h http.Handler
}

// 实现http.Handler接口
func (log *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 输出日志信息
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	// 使用下一个处理者处理请求
	log.h.ServeHTTP(w, r)
}
