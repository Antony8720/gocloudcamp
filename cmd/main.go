package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"gocloudcamp/internal/playlist"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	p := playlist.NewPlaylist("test")
	p.AddSong("test1", 15)
	p.AddSong("test2", 2)
	p.AddSong("test3", 7)
	p.AddSong("test4", 2)
	p.Play()
	time.Sleep(time.Second * 7)
	p.Stop()
	time.Sleep(time.Second * 3)
	p.Play()
	time.Sleep(time.Second * 5)
	p.Pause()
	time.Sleep(time.Second * 3)
	p.Play()
	<-ctx.Done()
}
