package main

import (
	"github.com/gin-gonic/gin"
//	"net/http"
    "fmt"
)


func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()

    router.POST("/post", func(c *gin.Context) {

        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")

        fmt.Printf("id: %s; page: %s; name: %s; message: %s\n\n", id, page, name, message)
    })
    router.Run(":8080")
}