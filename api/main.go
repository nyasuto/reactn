package main

import (
	"github.com/gin-gonic/gin"
)

type Index struct {
	Message string
}

var unusedValue = 10

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		index := Index{Message: "Hello world"}
		ctx.JSON(200, index)
		//ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
