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

// Start starts the timing.
func Start(name string) {
	defaultTimers.Start(name)
}

// Measure measures the elapsed time, and pause the timing.
func Measure(name string) time.Duration {
	return defaultTimers.Measure(name)
}

// MeasureCumulative measures the elapsed time, and do not pause the timing.
func MeasureCumulative(name string) time.Duration {
	return defaultTimers.MeasureCumulative(name)
}

// Pause pauses the timing.
func Pause(name string) {
	defaultTimers.Pause(name)
}

// Resume resumes the timing.
func Resume(name string) {
	defaultTimers.Resume(name)
}

// Start starts the timing.
func (t *Timers) Start(name string) {
	sw := t.safeGetSw(name)
	if sw == nil {
		sw = NewStopWatch(true)
		t.safeSetSw(name, sw)
	}
	sw.Start()
}

// Measure measures the elapsed time, and pause the timing.
func (t *Timers) Measure(name string) time.Duration {
	sw := t.safeGetSw(name)
	if sw == nil {
		sw = NewStopWatch(true)
		t.safeSetSw(name, sw)
	}
	sw.Pause()
	return sw.Elapsed()
}

// MeasureCumulative measures the elapsed time, and do not pause the timing.
func (t *Timers) MeasureCumulative(name string) time.Duration {
	sw := t.safeGetSw(name)
	if sw == nil {
		sw = NewStopWatch(true)
		t.safeSetSw(name, sw)
	}
	return sw.Elapsed()
}

// Pause pauses the timing.
func (t *Timers) Pause(name string) {
	sw := t.safeGetSw(name)
	if sw != nil {
		sw.Pause()
	}
}

// Resume resumes the timing.
func (t *Timers) Resume(name string) {
	sw := t.safeGetSw(name)
	if sw != nil {
		sw.Start()
	}
}

// Message returns the formatted elapsed time in milliseconds.
func (t *Timers) Message(name string) string {
	ms := float64(t.Measure(name).Milliseconds()) / 1000.0
	return fmt.Sprintf("%-8.3f ms %s", ms, name)
}

// safeGetNames returns the names of the timing StopWatch map.
func (t *Timers) safeGetNames() []string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return maps.Keys(t.stopWatches)
}

// safeSetSw sets the StopWatch in the timing map.
func (t *Timers) safeSetSw(name string, sw *StopWatch) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.stopWatches[name] = sw
}

// safeGetSw returns the StopWatch from the timing map.
func (t *Timers) safeGetSw(name string) *StopWatch {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.stopWatches[name]
}
