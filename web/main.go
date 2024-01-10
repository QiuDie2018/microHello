package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microHello/web/handler"
	"net/http"
)

const addr = ":8888"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "index",
	})
}

func main() {
	r := gin.Default()
	r.Handle("GET", "/", Index)
	hello := handler.HandleHello
	r.Handle("GET", "/hello", hello)
	if err := r.Run(addr); err != nil {
		fmt.Println("err")
	}
}
