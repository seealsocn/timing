package timing

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/maps"
)

var defaultTimers = NewTimers("defaultTimers")

// Timers measures the elapsed time.
type Timers struct {
	label       string
	stopWatches map[string]*StopWatch
	lock        sync.Mutex
}

// NewTimers creates a new Timers.
func NewTimers(label string) *Timers {
	return &Timers{
		label:       label,
		stopWatches: make(map[string]*StopWatch),
	}
}

// GetTimers returns the default Timers.
func GetTimers() *Timers {
	return defaultTimers
}

// Start starts the timer.
func Start(names ...string) {
	defaultTimers.Start(names...)
}

// Measure measures the elapsed time, and pauses the timer.
func Measure(name string) time.Duration {
	return defaultTimers.Measure(name)
}

// MeasureAll measures all the timers.
func MeasureAll() map[string]time.Duration {
	return defaultTimers.MeasureAll()
}

// Elapsed returns the elapsed time, and does not pause the timer.
func Elapsed(name string) time.Duration {
	return defaultTimers.Elapsed(name)
}

// ElapsedAll returns the elapsed time for all timers, and does not pause them.
func ElapsedAll() map[string]time.Duration {
	return defaultTimers.ElapsedAll()
}

// Pause pauses the timer.
func Pause(names ...string) {
	defaultTimers.Pause(names...)
}

// PauseAll pauses all the timers.
func PauseAll() {
	defaultTimers.PauseAll()
}

// Resume resumes the timer.
func Resume(names ...string) {
	defaultTimers.Resume(names...)
}

// Start starts the timer.
func (t *Timers) Start(names ...string) {
	at := time.Now()
	for _, name := range names {
		sw := t.safeGetSw(name)
		if sw == nil {
			sw = NewStopWatch(true)
			t.safeSetSw(name, sw)
		}
		sw.StartAt(at)
	}
}

// Measure measures the elapsed time, and pauses the timer.
func (t *Timers) Measure(name string) time.Duration {
	sw := t.safeGetSw(name)
	if sw == nil {
		sw = NewStopWatch(true)
		t.safeSetSw(name, sw)
	}
	sw.Pause()
	return sw.Elapsed()
}

// MeasureAll measures the elapsed time for all timers, and pauses them.
func (t *Timers) MeasureAll() map[string]time.Duration {
	elapsed := make(map[string]time.Duration)
	at := time.Now()
	for _, name := range t.safeGetNames() {
		sw := t.safeGetSw(name)
		if sw != nil {
			sw.PauseAt(at)
			elapsed[name] = sw.Elapsed()
		}
	}
	return elapsed
}

// Elapsed returns the elapsed time, and does not pause the timer.
func (t *Timers) Elapsed(name string) time.Duration {
	sw := t.safeGetSw(name)
	if sw == nil {
		sw = NewStopWatch(true)
		t.safeSetSw(name, sw)
	}
	return sw.Elapsed()
}

// ElapsedAll returns the elapsed time for all timers, and does not pause them.
func (t *Timers) ElapsedAll() map[string]time.Duration {
	elapsed := make(map[string]time.Duration)
	for _, name := range t.safeGetNames() {
		sw := t.safeGetSw(name)
		if sw != nil {
			elapsed[name] = sw.Elapsed()
		}
	}
	return elapsed
}

// Pause pauses the timer.
func (t *Timers) Pause(names ...string) {
	at := time.Now()
	for _, name := range names {
		sw := t.safeGetSw(name)
		if sw != nil {
			sw.PauseAt(at)
		}
	}
}

// PauseAll pauses all the timers.
func (t *Timers) PauseAll() {
	at := time.Now()
	for _, name := range t.safeGetNames() {
		sw := t.safeGetSw(name)
		if sw != nil {
			sw.PauseAt(at)
		}
	}
}

// Resume resumes the timer.
func (t *Timers) Resume(names ...string) {
	at := time.Now()
	for _, name := range names {
		sw := t.safeGetSw(name)
		if sw != nil {
			sw.StartAt(at)
		}
	}
}

// Message returns the formatted elapsed time in milliseconds.
func (t *Timers) Message(name string) string {
	ms := float64(t.Measure(name).Milliseconds()) / 1000.0
	return fmt.Sprintf("%-8.3f ms %s", ms, name)
}

// safeGetNames returns the names of the timer map.
func (t *Timers) safeGetNames() []string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return maps.Keys(t.stopWatches)
}

// safeSetSw sets the StopWatch in the timer map.
func (t *Timers) safeSetSw(name string, sw *StopWatch) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.stopWatches[name] = sw
}

// safeGetSw returns the StopWatch from the timer map.
func (t *Timers) safeGetSw(name string) *StopWatch {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.stopWatches[name]
}
