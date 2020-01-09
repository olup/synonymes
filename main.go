package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dictionary := open()

	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/:word", func(c *gin.Context) {
		word := c.Params.ByName("word")
		c.JSON(200, dictionary[word])
	})

	server.Run(":5000")

}
