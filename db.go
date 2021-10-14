package main

import (
	"encoding/json"
	"os"
)

type DB struct {
	path   string
	videos []*Video
}

func NewDB(path string) *DB {
	return &DB{path: path}
}

func (d *DB) Init() {
	d.loadFromDisc()
}

func (d *DB) Close() {
	d.saveToDisc()
}

func (d *DB) saveToDisc() error {
	bytes, err := json.Marshal(d.videos)
	if err != nil {
		return err
	}

	err = os.WriteFile(d.path, bytes, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

func (d *DB) loadFromDisc() error {
	if _, err := os.Stat(d.path); os.IsNotExist(err) {
		d.videos = []*Video{}
		return nil
	}

	bytes, err := os.ReadFile(d.path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &d.videos)
	if err != nil {
		panic(err)
	}

	return nil
}

func (d *DB) AddVideo(video *Video) {
	d.videos = append(d.videos, video)
	if err := d.saveToDisc(); err != nil {
		panic(err)
	}
}

func (d *DB) AddVideos(videos []*Video) {
	d.videos = append(d.videos, videos...)

	if err := d.saveToDisc(); err != nil {
		panic(err)
	}
}
