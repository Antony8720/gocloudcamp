package playlist

import (
	"fmt"
	"time"
)

type event int

const (
	eventPlay    event = iota // 0
	eventStop                 // 1
	eventNext                 // 2
	eventPrev                 //3
	eventPause                //4
	eventPlaying              // 5
)

type Song struct {
	name     string
	duration int
	previous *Song
	next     *Song
}

type Playlist struct {
	name       string
	head       *Song
	tail       *Song
	nowPlaying *Song
	pause      bool
	// next       bool
	// prev       bool

	playingTime time.Duration
	events      chan event
}

func NewPlaylist(name string) *Playlist {
	p := &Playlist{
		name:   name,
		events: make(chan event, 1),
	}

	go p.processEventLoop()

	return p
}

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
		case eventNext:
			if p.nowPlaying != nil {
				p.nowPlaying = p.nowPlaying.next
			} else {
				p.nowPlaying = p.head
			}
			p.playingTime = 0
		case eventPrev:
			if p.nowPlaying != nil {
				p.nowPlaying = p.nowPlaying.previous
			}
			p.playingTime = 0
		case eventPause:
			p.pause = true
		default:
		}

		p.streamData()
	}
}

func (p *Playlist) streamData() {

	const playingPart = 5 * time.Second
	if p.nowPlaying == nil || p.pause == true {
		return
	}

	fmt.Println("dur", p.nowPlaying.duration, "playing", p.playingTime.Seconds())
	diff := min(int(playingPart.Seconds()), p.nowPlaying.duration-int(p.playingTime.Seconds()))
	fmt.Println("play", p.nowPlaying.name, diff, "sec")
	time.Sleep(time.Duration(diff) * time.Second)

	p.playingTime += time.Duration(diff) * time.Second
	if int(p.playingTime.Seconds()) >= p.nowPlaying.duration {
		fmt.Println("end playing", p.nowPlaying.name)
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

func min(val1, val2 int) int {
	if val1 <= val2 {
		return val1
	}

	return val2
}

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

func (p *Playlist) Play() {
	p.sendEvent(eventPlay)
}

func (p *Playlist) Pause() {
	fmt.Println("pause")
	p.sendEvent(eventPause)
}

func (p *Playlist) AddSong(name string, duration int) {
	s := &Song{
		name:     name,
		duration: duration,
	}
	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	p.tail = s
}

func (p *Playlist) Next() {
	p.sendEvent(eventNext)
}

func (p *Playlist) Prev() {
	p.sendEvent(eventPrev)
}

func (p *Playlist) Stop() {
	fmt.Println("Stop")
	p.sendEvent(eventStop)
}
