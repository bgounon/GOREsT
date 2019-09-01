package main

func buildURLFromVideoID(videoID string) string {

	return youtubeBaseURL + "/watch?v=" + videoID
}

func buildEmbedCodeFromVideoID(key int, videoID string) string {
	return "<iframe type='text/html' width='640'height='360' src=" + youtubeBaseURL + "/embed/" + videoID + "?html5=1 frameborder='0'></iframe><br><br>"
}
