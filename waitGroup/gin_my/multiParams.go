package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()

    // This handler will match /user/john but will not match neither /user/ or /user
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    // However, this one will match /user/john/ and also /user/john/send
    // If no other routers match /user/john, it will redirect to /user/john/
    router.GET("/user/:name/:action/*abcd", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        abcd := c.Param("abcd")
        message := name + " is " + action + " and " + abcd
        c.String(http.StatusOK, message)
    })

    router.Run(":8080")
}