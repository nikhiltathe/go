package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "fmt"
)


func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()

   // Query string parameters are parsed using the existing underlying request object.
    // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
    router.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        //lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
        lastname := c.Request.URL.Query().Get("lastname")

        fmt.Println("firstname :",firstname)
        fmt.Println("lastname :",lastname)
        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })
    
    router.Run(":8080")
}