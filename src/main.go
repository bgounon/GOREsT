package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	youtubeService := youtubeInit()

	r := gin.Default()
	r.GET("/playlist/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, getTitlesFromPlaylistID(id, youtubeService))
	})

	r.GET("/playlist/:id/embed", func(c *gin.Context) {
		id := c.Param("id")
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, embedAllVideos(id, youtubeService))
	})

	r.Run()
}
