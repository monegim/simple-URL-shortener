package main

import (
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	
	r := gin.Default()
	r.GET("/:catch_all", mainHandler)
	r.Run(":8081")
}

func mainHandler(ctx *gin.Context) {
	log.Println("Main is called")
	url := ctx.Request.URL
	log.Println(url, "is hit")
}