package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Index struct {
	Message string
	Files []string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		//index := Index{Message: "Hello world"}
		//ctx.JSON(200, index)
		ctx.HTML(200, "index.html", gin.H{})		
	})

	router.GET("/movie", func(ctx *gin.Context) {
		files, err := ioutil.ReadDir("/mnt")

		if err != nil{
			ctx.JSON(200, Index{Message: err.Error()})
		}
		var result []string

		for _, f := range files{
			result = append(result, f.Name())
			// result = append(result, f.Name)
		}
		ctx.JSON(200, Index{Files: result})

		//index := Index{Message: "Hello world"}
		//ctx.JSON(200, index)
	})

	router.Run()
}
