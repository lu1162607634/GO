package main

import (
	"github.com/gin-gonic/gin"
	"httpcache/lusilu/Class03/router"
)

func main() {
	mux := gin.Default()
	router.New().Install(mux)
	mux.Run(":8082")
}
