package playlist

import (
	"fmt"
	"log"
	"time"
)

type event int

const (
	eventPlay    event = iota // 0
	eventStop                 // 1
	eventNext                 // 2
	eventPrev                 // 3
	eventPause                // 4
	eventPlaying              // 5
)

type Song struct {
	name     string
	duration int
	previous *Song
	Next     *Song
}

type Playlist struct {
	head       *Song
	tail       *Song
	nowPlaying *Song
	pause      bool

	playingTime time.Duration
	events      chan event
}

// Creating new playlist
func NewPlaylist() *Playlist {
	p := &Playlist{
		events: make(chan event, 1),
	}

	go p.processEventLoop()

	return p
}

// Event processing
func (p *Playlist) processEventLoop() {
	for {
		event := <-p.events
		switch event {
		case eventPlay:
			p.pause = false
			if p.nowPlaying == nil {
				p.nowPlaying = p.head
			}
		case eventStop:
			p.nowPlaying = nil
			p.playingTime = 0
			p.head = nil
			p.tail = nil
		case eventNext:
			if p.nowPlaying != nil {
				if p.nowPlaying.Next != nil {
					p.nowPlaying = p.nowPlaying.Next
				} else {
					p.nowPlaying = p.head
				}
			} else {
				p.nowPlaying = p.head
			}
			p.playingTime = 0
		case eventPrev:
			if p.nowPlaying != nil {
				if p.nowPlaying.previous != nil {
					p.nowPlaying = p.nowPlaying.previous
				}
			}
			p.playingTime = 0
		case eventPause:
			p.pause = true
		default:
		}

		p.streamData()
	}
}

// Streaming data
func (p *Playlist) streamData() {

	const playingPart = 5 * time.Second
	if p.nowPlaying == nil || p.pause == true {
		return
	}

	diff := min(int(playingPart.Seconds()), p.nowPlaying.duration-int(p.playingTime.Seconds()))

	log.Println("play", p.nowPlaying.name, diff, "sec")
	time.Sleep(time.Duration(diff) * time.Second)

	p.playingTime += time.Duration(diff) * time.Second

	if int(p.playingTime.Seconds()) >= p.nowPlaying.duration {
		log.Println("end playing", p.nowPlaying.name)
		p.playingTime = 0
		select {
		case p.events <- eventNext:
		default:
		}
		return
	}
	select {
	case p.events <- eventPlaying:
	default:
	}
}

// Helper function min
func min(val1, val2 int) int {
	if val1 <= val2 {
		return val1
	}

	return val2
}

// Send event to channel
func (p *Playlist) sendEvent(e event) {
	for {
		select {
		case p.events <- e:
			return
		default:
			select {
			case <-p.events:
			default:
			}
		}
	}
}

// Play playlist
func (p *Playlist) Play() {
	log.Println("Play")
	p.sendEvent(eventPlay)
}

// Pause playlist
func (p *Playlist) Pause() {
	log.Println("Pause")
	p.sendEvent(eventPause)
}

// Add song to playlist
func (p *Playlist) AddSong(name string, duration int) {
	s := &Song{
		name:     name,
		duration: duration,
	}
	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.Next = s
		s.previous = p.tail
	}
	p.tail = s
}

// Switch to next song
func (p *Playlist) Next() {
	p.sendEvent(eventNext)
}

// Swith to previous song
func (p *Playlist) Prev() {
	p.sendEvent(eventPrev)
}

// Stop playing playlist
func (p *Playlist) Stop() {
	p.sendEvent(eventStop)
}

// Show now playing song
func (p *Playlist) ShowSong() (string, int) {
	return p.nowPlaying.name, p.nowPlaying.duration
}

type dbPlaylist struct {
	Name     string
	Duration int
}

func (p *Playlist) SetMemoryPlaylist() []dbPlaylist {
	song := p.head
	fmt.Println(song.name)
	playlist := make([]dbPlaylist, 0)
	for song != nil {
		playlistModel := &dbPlaylist{Name: song.name, Duration: song.duration}
		playlist = append(playlist, *playlistModel)
		song = song.Next
	}
	return playlist

}

func (p *Playlist) IsEmpty() bool {
	if p.head == nil {
		return true
	}
	return false
}

func (p *Playlist) IsPlaying() bool {
	if p.nowPlaying == nil {
		return false
	}
	return true
}
