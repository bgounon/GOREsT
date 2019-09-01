package main

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YTVideo : contains useful info about a given video
type YTVideo struct {
	URL   string
	Title string
}

var youtubeBaseURL string

func youtubeInit(apiKey string) *youtube.Service {

	youtubeBaseURL = "https://www.youtube.com"

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	return youtubeService
}
