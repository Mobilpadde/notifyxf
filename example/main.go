package main

import (
	"os"

	"github.com/Mobilpadde/notifyxf"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(notifyxf.Recover(os.Getenv("NOTIFYXF_TOKEN"), "My Back-end Server"))

	router.GET("/", func(c *gin.Context) {
		panic("hai")
	})

	router.Run()
}
