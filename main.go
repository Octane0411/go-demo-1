package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test1")
	e := gin.Default()
	e.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello"})
	})
	e.Run(":8080")
}
