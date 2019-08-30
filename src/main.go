package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	b, err := ioutil.ReadFile("../.APIKEY")
	if err != nil {
		fmt.Print(err)
	}
	apiKey := strings.TrimSuffix(string(b), "\n")

	r := gin.Default()
	r.GET("/playlist/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, getTitlesFromPlaylistID(id, apiKey))
	})
	r.Run()
}
