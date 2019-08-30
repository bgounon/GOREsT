package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YTVideo : contains useful info about a given video
type YTVideo struct {
	ID    string
	Title string
}

func getTitlesFromPlaylistID(playListID string, apiKey string) []YTVideo {

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Print(err)
	}

	call := youtubeService.PlaylistItems.List("snippet")
	call = call.PlaylistId(playListID)
	call = call.MaxResults(50)
	playList, err := call.Do()
	if err != nil {
		fmt.Print(err)
	}

	var videos []YTVideo

	for ok := true; ok; ok = !(playList.NextPageToken == "") {
		for _, item := range playList.Items {
			var video YTVideo
			video.ID = item.Id
			video.Title = item.Snippet.Title
			videos = append(videos, video)
		}

		playList, err = call.PageToken(playList.NextPageToken).Do()
		if err != nil {
			fmt.Print(err)
		}
	}

	return videos
}
