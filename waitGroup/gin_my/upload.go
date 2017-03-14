package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"
import "io"
import "log"

func main() {
    router := gin.Default()

    router.POST("/upload", func(c *gin.Context) {

            file, header , err := c.Request.FormFile("upload")
            filename := header.Filename
            fmt.Println(header.Filename)
            out, err := os.Create("./tmp/"+filename+".png")
            if err != nil {
                log.Fatal(err)
            }
            defer out.Close()
            _, err = io.Copy(out, file)
            if err != nil {
                log.Fatal(err)
            }   
    })
    router.Run(":8080")
}