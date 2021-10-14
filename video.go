package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Video struct {
	Id           int    `json:"id"`
	VideoId      string `json:"video_id"`
	Title        string `json:"title"`
	IsDownloaded bool   `json:"is_downloaded"`
	IsFetched    bool   `json:"is_fetched"`
	CreatedAt    string `json:"created_at"`
}

type noEmbed struct {
	Url          string `json:"url"`
	Version      string `json:"version"`
	Title        string `json:"title"`
	AuthorName   string `json:"author_name"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

func (v *Video) FetchVideo() {
	res, err := http.Get(fmt.Sprintf("https://noembed.com/embed?url=https://www.youtube.com/watch?v=%s", v.VideoId))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	embedRes := &noEmbed{}
	json.Unmarshal(body, &embedRes)
	v.Title = embedRes.Title
	v.IsFetched = true
}
