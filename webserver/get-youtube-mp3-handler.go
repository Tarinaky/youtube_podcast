package main

import (
	"log"
	"net/http"
	"youtube_podcast/youtube_mp3"

	"github.com/gorilla/mux"
)

func getYoutubeMP3Handler(w http.ResponseWriter, r *http.Request) {
	const chunkSize = 1024
	videoID := mux.Vars(r)["id"]
	log.Printf("DEBUG: Serving video %s", videoID)
	defer log.Printf("DEBUG: Finished serving %s", videoID)
	stream, err := youtube_mp3.NewAudioStream(videoID)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	buffer := make([]byte, chunkSize)
	for {
		_, err = stream.Read(buffer)
		if err != nil {
			break
		}

		_, err = w.Write(buffer)
		if err != nil {
			break
		}
	}
}
