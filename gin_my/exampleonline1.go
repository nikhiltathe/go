/*authenticate("user", "pass");

function authenticate(username, password) {
  $.ajax({
    type: "POST",
    url: "http://localhost:3000/login",
    contentType: "application/json",
    async: false,
    data: '{"username": "' + username + '", "password" : "' + password + '"}',
    success: function (res) {
      console.log(res);
    }
  });
}
*/
package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    g := gin.New()

    // Logging middleware
    g.Use(gin.Logger())

    // Recovery middleware
    g.Use(gin.Recovery())

    // CORS middleware
    g.Use(CORSMiddleware())

    g.POST("/login", LoginCheck)

    // Serve
    g.Run(":3000")
}

type LoginJSON struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func LoginCheck(c *gin.Context) {
    var json LoginJSON
    c.Bind(&json)
    c.JSON(200, json)
    // Print response
    fmt.Println(json.Username, json.Password)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.Abort()
            return
        }
        c.Next()
    }


}
