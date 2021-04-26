package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// original
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// #1
	r.GET("/friends", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"array": [4]int{1, 2, 3, 4},
		})
	})

	// #2
	r.GET("/name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi! I am Thomas.",
		})
	})

	// #3
	r.GET("/birthday", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"message": "",
		})
	})

	r.POST("/friends", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
	})

	r.Run()
}
