package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "log"
    "github.com/gocql/gocql"
)


 // Binding from JSON
type Login struct {
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

type Tweet struct {
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
    Text string `form:"text" json:"text" binding:"required"`
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

   // Example for binding JSON ({"user": "manu", "password": "123"})
    /* Before you execute the program, Launch `cqlsh` and execute:
    create keyspace example1 with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
    create table example1.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
    create index on example1.tweet(timeline);
    */

    router.POST("/addJSON", func(c *gin.Context) {
        var json Tweet
        if c.BindJSON(&json) == nil {
            if json.User == "manu" && json.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})


                cluster := gocql.NewCluster("127.0.0.1")
                cluster.Keyspace = "example1"
                cluster.Consistency = gocql.Quorum
                session, _ := cluster.CreateSession()
                defer session.Close()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               

               // insert a tweet
                if err := session.Query("INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)",
                    json.User , gocql.TimeUUID(), json.Text).Exec(); err != nil {
                    log.Fatal(err)
                }
    
                var id gocql.UUID
                var text string

                /* Search for a specific set of records whose 'timeline' column matches
                * the value 'me'. The secondary index that we created earlier will be
                * used for optimizing the search */
                if err := session.Query("SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1",
                    json.User).Consistency(gocql.One).Scan(&id, &text); err != nil {
                    log.Fatal(err)
                }
                fmt.Println("Tweet:", id, text)

                // list all tweets
                iter := session.Query("SELECT id, text FROM tweet WHERE timeline = ?", json.User).Iter()
                for iter.Scan(&id, &text) {
                    fmt.Println("Tweet:", id, text)
                }
                if err := iter.Close(); err != nil {
                    log.Fatal(err)
                }
            }
        }    
    })


    // Listen and server on 0.0.0.0:8080
    router.Run(":8080")

/*
    curl -H "Content-Type: application/json" -X GET -d '{"user":"manu","password":"123"}' http://localhost:8080/getJSON
    curl -H "Content-Type: application/json" -X POST -d '{"user":"manu","password":"123"}' http://localhost:8080/loginJSON
    curl -H "Content-Type: application/json" -X POST -d '{"user":"manu","password":"123", "text":"Package text"}' http://localhost:8080/addJSON

*/
}