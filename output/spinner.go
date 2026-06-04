// Package output
package output

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var frames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// Tracker tracks which tools are currently running
type Tracker struct {
	mu      sync.Mutex
	running []string
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) Add(tool string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.running = append(t.running, tool)
}

func (t *Tracker) Remove(tool string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i, v := range t.running {
		if v == tool {
			t.running = append(t.running[:i], t.running[i+1:]...)
			return
		}
	}
}

// Print clears the spinner line, prints the message, then lets the spinner redraw
func (t *Tracker) Print(msg string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	fmt.Printf("\r\033[K%s\n", msg)
}

func (t *Tracker) status() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.running) == 0 {
		return "Finishing up..."
	}
	return "Running: [" + strings.Join(t.running, ", ") + "]"
}

// Spinner displays currently running tools until done is signaled
func Spinner(tracker *Tracker, done chan bool) {
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("\r\033[K")
			return
		default:
			fmt.Printf("\r %s  %s\033[K", CyanColor(frames[i%len(frames)]), tracker.status())
			time.Sleep(80 * time.Millisecond)
			i++
		}
	}
}
