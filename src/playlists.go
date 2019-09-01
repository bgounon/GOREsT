package main

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func getTitlesFromPlaylistID(playListID string, youtubeService *youtube.Service) []YTVideo {
	fmt.Println(youtubeService)
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
			video.Title = item.Snippet.Title
			video.URL = buildURLFromVideoID(item.Snippet.ResourceId.VideoId)
			videos = append(videos, video)
		}

		playList, err = call.PageToken(playList.NextPageToken).Do()
		if err != nil {
			fmt.Print(err)
		}
	}

	return videos
}
