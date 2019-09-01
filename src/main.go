package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var apiKey string
	flag.StringVar(&apiKey, "api", "", "Youtube API Key")
	flag.Parse()
	youtubeService := youtubeInit(apiKey)

	r := gin.Default()
	r.GET("/playlist/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, getTitlesFromPlaylistID(id, youtubeService))
	})

	r.Run()
}
