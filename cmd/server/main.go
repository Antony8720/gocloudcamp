package main

import (
	"gocloudcamp/internal/api/server"
	"gocloudcamp/internal/playlist"
)

func main() {
	p := playlist.NewPlaylist()
	s := server.New(p)
	s.Start(s)
}
