package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		print("Hello!!!")
	})

	fmt.Sprintf("🚀   CI/CD is working on Windows runner!\n")

	if err := g.Run(":8080"); err != nil {
		panic("Exception occurred !!!")
	}

}
