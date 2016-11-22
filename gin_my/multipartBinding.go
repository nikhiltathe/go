package main

import (
    "github.com/gin-gonic/gin"
//    "github.com/gin-gonic/gin/binding"
)

type LoginForm struct {
    User     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func main() {
    router := gin.Default()
    router.POST("/login", func(c *gin.Context) {
        // you can bind multipart form with explicit binding declaration:
        // c.BindWith(&form, binding.Form)
        // or you can simply use autobinding with Bind method:
        var form LoginForm
        // in this case proper binding will be automatically selected
        if c.Bind(&form) == nil {
            if form.User == "user" && form.Password == "password" {
                c.JSON(200, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(401, gin.H{"status": "unauthorized"})
            }
        }
    })
    router.Run(":8080")
}