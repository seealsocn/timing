package timing

import (
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	sw := NewStopWatch(true)
	time.Sleep(2 * time.Millisecond)
	if sw.Elapsed().Milliseconds() == 0 {
		t.Errorf("elapsed not working")
	}
}

func TestStart(t *testing.T) {
	sw := NewStopWatch(false)

	time.Sleep(2 * time.Millisecond)
	ms1 := sw.Elapsed().Milliseconds()
	if ms1 != 0 {
		t.Errorf("autoStart=false not working, ms1=%d", ms1)
	}

	sw.Start()
	time.Sleep(2 * time.Millisecond)
	ms2 := sw.Elapsed().Milliseconds()
	if ms2 < 1 {
		t.Errorf("start not working")
	}
}

func TestPause(t *testing.T) {
	sw := NewStopWatch(true)
	time.Sleep(2 * time.Millisecond)
	sw.Pause()

	elapsed1 := sw.Elapsed()
	time.Sleep(2 * time.Millisecond)
	elapsed2 := sw.Elapsed()

	if elapsed1 != elapsed2 {
		t.Errorf("pause not working")
	}
}

func TestStop(t *testing.T) {
	sw := NewStopWatch(true)
	time.Sleep(2 * time.Millisecond)
	sw.Stop()

	elapsed1 := sw.Elapsed()
	time.Sleep(2 * time.Millisecond)
	elapsed2 := sw.Elapsed()

	if elapsed1 != elapsed2 {
		t.Errorf("stop not working")
	}
}

func TestRestart(t *testing.T) {
	sw := NewStopWatch(true)
	time.Sleep(2 * time.Millisecond)
	elapsed1 := sw.Elapsed()

	sw.Restart()
	elapsed2 := sw.Elapsed()

	if elapsed1 <= elapsed2 {
		t.Errorf("restart not working")
	}
}
