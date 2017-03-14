package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

 // Binding from JSON
type Login struct {
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

func main() {
    router := gin.Default()

    // --------------------------------------------------------------------------
    // Trial part
    // --------------------------------------------------------------------------
    // Example for binding JSON ({"user": "manu", "password": "123"})

    router.GET("/getJSON", func(c *gin.Context) {
        var json Login
        if c.BindJSON(&json) == nil {
            if json.User == "manu" && json.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
        fmt.Println("json is ",json)
    })

    // --------------------------------------------------------------------------

    // Example for binding JSON ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if c.BindJSON(&json) == nil {
            if json.User == "manu" && json.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
    })

    // Example for binding a HTML form (user=manu&password=123)
    // Content-Type: application/x-www-form-urlencoded
    router.POST("/loginForm", func(c *gin.Context) {
        var form Login
        // This will infer what binder to use depending on the content-type header.
        if c.Bind(&form) == nil {
            if form.User == "manu" && form.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
    })

    // Listen and server on 0.0.0.0:8080
    router.Run(":8080")

/*
    curl -H "Content-Type: application/json" -X POST -d '{"user":"manu","password":"123"}' http://localhost:8080/loginJSON
*/
}