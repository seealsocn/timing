package timing

import "time"

// StopWatch measures the elapsed time.
type StopWatch struct {
	start   *time.Time
	end     *time.Time
	elapsed time.Duration
}

// NewStopWatch creates a StopWatch.
func NewStopWatch(autoStart bool) *StopWatch {
	var start *time.Time
	if autoStart {
		now := time.Now()
		start = &now
	}
	return &StopWatch{
		start: start,
	}
}

// Start starts the stopwatch. It sets the `start` time to the current time.
func (s *StopWatch) Start() {
	now := time.Now()
	s.start = &now
	s.end = nil
}

// StartAt starts the stopwatch. It sets the `start` time to the specified time. Useful for testing.
func (s *StopWatch) StartAt(at time.Time) {
	s.start = &at
	s.end = nil
}

// Restart restarts the stopwatch. It sets `start` time to the current time, and resets the `end` time and `elapsed` time.
func (s *StopWatch) Restart() {
	now := time.Now()
	s.start = &now
	s.end = nil
	s.elapsed = time.Duration(0)
}

// RestartAt restarts the stopwatch. It sets `start` time to the specified time, and resets the `end` time and `elapsed` time. Useful for testing.
func (s *StopWatch) RestartAt(at time.Time) {
	s.start = &at
	s.end = nil
	s.elapsed = time.Duration(0)
}

// Pause adds the current elapsed time to `elapsed` and resets the `start` time.
func (s *StopWatch) Pause() {
	if s.start == nil {
		return
	}
	if s.end == nil {
		s.elapsed += time.Since(*s.start)
	} else {
		s.elapsed += s.end.Sub(*s.start)
	}
	s.start = nil
}

// PauseAt adds the specified elapsed time to `elapsed` and resets the `start` time. Useful for testing.
func (s *StopWatch) PauseAt(at time.Time) {
	if s.start == nil {
		return
	}
	if s.end == nil {
		s.elapsed += at.Sub(*s.start)
	} else {
		s.elapsed += s.end.Sub(*s.start)
	}
	s.start = nil
}

// Stop stops the timer. It sets the `end` time to the current time.
func (s *StopWatch) Stop() {
	now := time.Now()
	s.end = &now
}

// StopAt stops the timer. It sets the `end` time to the specified time. Useful for testing.
func (s *StopWatch) StopAt(at time.Time) {
	s.end = &at
}

// Elapsed returns the Duration since the last start.
func (s *StopWatch) Elapsed() time.Duration {
	if s.start == nil {
		return s.elapsed
	}
	if s.end == nil {
		return time.Since(*s.start) + s.elapsed
	}
	return s.end.Sub(*s.start) + s.elapsed
}
