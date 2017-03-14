package main

import (
	"github.com/gin-gonic/gin"
//	"net/http"
//    "fmt"
)


func main() {
        // Creates a router without any middleware by default
    r := gin.New()
    gin.SetMode(gin.DebugMode)

    // Global middleware
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Per route middleware, you can add as many as you desire.
    r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

    // Authorization group
    // authorized := r.Group("/", AuthRequired())
    // exactly the same than:
    authorized := r.Group("/")
    // per group middleware! in this case we use the custom created
    // AuthRequired() middleware just in the "authorized" group.
    authorized.Use(AuthRequired())
    {
        authorized.POST("/login", loginEndpoint)
        authorized.POST("/submit", submitEndpoint)
        authorized.POST("/read", readEndpoint)

        // nested group
        testing := authorized.Group("testing")
        testing.GET("/analytics", analyticsEndpoint)
    }

    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")
}