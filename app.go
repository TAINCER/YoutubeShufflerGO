package main

import (
	"context"
)

// App struct
type App struct {
	ctx            context.Context
	bookmarkParser *BookmarkParser
}

// NewApp creates a new App application struct
func NewApp(bookmarkParser *BookmarkParser) *App {
	return &App{bookmarkParser: bookmarkParser}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
	b.bookmarkParser.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}
