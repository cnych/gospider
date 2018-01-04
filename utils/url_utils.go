package utils

import (
	"fmt"
	"regexp"
)

func Match1(pattern, text string) string {
	r, _ := regexp.Compile(pattern)
	return r.FindString(text)
}

func UrlToModule(url string) {
	videoHost := Match1("https?://([^/]+)/", url)
	videoUrl := Match1("https?://[^/]+(.*)", url)
	fmt.Printf("video_host=%s, video_url=%s\n", videoHost, videoUrl)
}
