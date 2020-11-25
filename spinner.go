package spinner

import (
	"fmt"
	"sync"
	"time"
)

var Dots = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`

// Spinner type
type Spinner struct {
	text     string // Text to display before the spinner.
	mu       sync.Mutex
	frames   []rune
	length   int
	pos      int
	stopChan chan struct{}
	spinning bool
}

// New returns spinner instance
func New(text string) *Spinner {
	frames := []rune(Dots)
	s := &Spinner{
		text:     text,
		frames:   frames,
		length:   len(frames),
		stopChan: make(chan struct{}, 1),
	}
	return s
}

// Start the spinner.
func (s *Spinner) Start() {
	if s.spinning {
		return
	}

	s.mu.Lock()
	s.spinning = true
	s.mu.Unlock()

	go func() {
		for {
			select {
			case <-s.stopChan:
				fmt.Print("\r\033[K")
				return
			default:
				fmt.Printf("\r%s %s", s.Next(), s.text)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

// Stop the spinner.
func (s *Spinner) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.spinning {
		return
	}

	s.stopChan <- struct{}{}
	s.spinning = false
}

// Next returns next rune in Dots
func (s *Spinner) Next() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.frames[s.pos%s.length]
	s.pos++
	return string(r)
}
