package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	_ "leetcode/打包html进二进制/statik" // TODO: Replace with the absolute import path
	"net/http"
)

// statik -src=./html2

func main() {
	r := gin.Default()

	statikFS, err := fs.New()
	if err != nil {
		fmt.Println(err)
	}

	r.GET("/", gin.WrapH(http.FileServer(statikFS)))

	r.Run(":9080")
}
