package main

import (
	"github.com/gin-gonic/gin"
//	"net/http"
    "fmt"
)

func main() {
    router := gin.Default()

    router.POST("/post", func(c *gin.Context) {

        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")

        fmt.Printf("\nid: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
    })
    router.Run(":8080")

    /* Query
    http://localhost:8080/post?id=1234&page=1
    */

    /* Request header
    Content-Type: application/x-www-form-urlencoded
    */

    /* Request body
    name=manu&message=this_is_great
    */
}