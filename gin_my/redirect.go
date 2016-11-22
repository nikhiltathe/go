package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
//    "fmt"
)


func main() {
    router := gin.Default()
    router.GET("/test", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
    })
    router.Run(":8080")
}