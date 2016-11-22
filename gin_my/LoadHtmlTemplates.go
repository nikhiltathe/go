package main

import "github.com/gin-gonic/gin"
//import "fmt"

func main() {
	r := gin.Default()

	r.LoadHTMLTemplates("templates/*")
	r.GET("/index", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tmpl", obj)
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}