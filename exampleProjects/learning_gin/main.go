package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	err := InitGinServer(os.Stdout)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

}

func InitGinServer(writer io.Writer) (err error) {
	if writer == nil {
		writer = os.Stdout
	}

	//create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	//create an endpoint
	r.GET("/ping", func(c *gin.Context) {
		// c is the passed context
		// TODO: Gin contexts ...

		c.JSON(http.StatusOK, gin.H{ // TODO: what happening here
			"message": "pong",
		})
	})

	r.SetTrustedProxies(nil) //Gin will look for and believe forwarded IP headers from trusted proxy addresses
	//otherwise, if a peer is not trusted, peer's IP will be used
	//start a server with default options on default port 8080: r.Run()
	err = r.Run("127.0.0.1:8080")

	return

}
