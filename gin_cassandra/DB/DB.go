package db

import (
    "github.com/gin-gonic/gin"
    "net/http"
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

func Add(c *gin.Context,json Tweet)() {

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
        } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        }
    }
}
/*
func read(c *gin.Context,json Tweet)() {

        if c.BindJSON(&json) == nil {
        if json.User == "manu" && json.Password == "123" {
            c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})


            cluster := gocql.NewCluster("127.0.0.1")
            cluster.Keyspace = "example1"
            cluster.Consistency = gocql.Quorum
            session, _ := cluster.CreateSession()
            defer session.Close()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               

            var id gocql.UUID
            var text string

                // Search for a specific set of records whose 'timeline' column matches
                // the value 'me'. The secondary index that we created earlier will be
                // used for optimizing the search 
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

            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
    }
*/
/*
    curl -H "Content-Type: application/json" -X POST -d '{"user":"manu","password":"123"}' http://localhost:8080/loginJSON
*/
