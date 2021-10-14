package main

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type BookmarkParser struct {
	ctx context.Context
	db  *DB
}

func NewBookmarkParser(db *DB) *BookmarkParser {
	return &BookmarkParser{db: db}
}

func (b *BookmarkParser) ParseFromFile() []*Video {
	fileName, err := runtime.OpenFileDialog(b.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		panic(err)
	}

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	videos := []*Video{}

	re := regexp.MustCompile(`https:\/\/www\.youtube\.com\/watch\?v=([-a-zA-Z0-9_!"$%]){11}`)
	for _, match := range re.FindAllString(string(bytes), -1) {
		// TODO: Also include shorthanded YouTube URLs
		// And also check if more than one GET Params is present, wheather its a video
		// from an Playlist or the timestamp is in the title, this may be a problem when
		// preloading or downloading the video
		splitMatch := strings.Split(match, "?v=")
		videos = append(videos, &Video{VideoId: splitMatch[len(splitMatch)-1]})
	}

	b.db.AddVideos(videos)

	return nil
}
