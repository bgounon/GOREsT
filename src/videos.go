package main

func buildURLFromVideoID(videoID string) string {

	return youtubeBaseURL + "/watch?v=" + videoID
}
