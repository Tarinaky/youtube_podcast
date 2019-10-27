package youtube_rss

import (
	"bufio"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/feeds"
)

const timestampFormat = "20060102"

func GetFeed(path string) (*feeds.Feed, error) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:   "Placeholder: "+path,
		Link:    &feeds.Link{Href: "http://youtube.com/" + path},
		Author:  &feeds.Author{Name: "youtube_podcast"},
		Created: now,
	}

	source, err := NewJSONStream(path)
	if err != nil {
		return feed, err
	}
	defer source.Close()

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		buffer := scanner.Bytes()
		var video Video
		json.Unmarshal(buffer, &video)

		time, err := time.Parse(timestampFormat, video.UploadDate)
		if err != nil {
			log.Printf("ERROR: GetFeed: %s", err)
		}
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       video.Title,
			Link:        &feeds.Link{Href: "127.0.0.1:8080/youtube/mp3/" + video.VideoID},
			Description: video.Description,
			Author:      &feeds.Author{Name: guessAuthor(video)},
			Created:     time,
		})
	}

	return feed, nil
}

func guessAuthor(video Video) string {
	if video.Creator != "" {
		return video.Creator
	}
	return video.Uploader
}
