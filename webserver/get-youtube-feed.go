package main

import (
	"log"
	"net/http"
	"strings"
	"youtube_podcast/youtube_rss"
)

func getYoutubeFeed(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/youtube/feed/") + "?" + r.URL.RawQuery
	log.Printf("DEBUG: Fetching feed for %s", path)
	feed, err := youtube_rss.GetFeed(path)
	if err != nil {
		log.Panic(err)
	}
	rss, err := feed.ToRss()
	if err != nil {
		log.Panic(err)
	}
	w.Write([]byte(rss))
}
