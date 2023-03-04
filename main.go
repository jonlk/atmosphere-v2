package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/atmosphere-v2/web"
)

var (
	mode string = gin.DebugMode
)

func main() {

	gin.SetMode(mode)

	router := gin.Default()
	router.SetTrustedProxies([]string{""})

	web.RegisterRoutes(router.Group("/api"))

	log.Fatal(router.Run(":3000"))
}
