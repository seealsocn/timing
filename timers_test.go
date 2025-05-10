package timing

import (
	"testing"
	"time"
)

func TestMeasure(t *testing.T) {
	tm := GetTimers()

	tm.Start("m1")
	time.Sleep(1 * time.Millisecond)
	d1 := tm.Measure("m1")

	tm.Start("m2")
	time.Sleep(2 * time.Millisecond)
	d2 := tm.Measure("m2")

	if d1 >= d2 {
		t.Errorf("measure not working")
	}

	if d1 != tm.Measure("m1") {
		t.Errorf("measure pause not working")
	}

	Start("m3")
	time.Sleep(2 * time.Millisecond)
	d3 := Measure("m3")
	if d3.Milliseconds() < 1 {
		t.Errorf("timing measure not working")
	}
}

func TestTimingElapsed(t *testing.T) {
	tm := GetTimers()

	tm.Start("m1")
	time.Sleep(2 * time.Millisecond)
	d1 := tm.Elapsed("m1")

	time.Sleep(1 * time.Millisecond)
	d2 := tm.Elapsed("m1")

	if d1 >= d2 {
		t.Errorf("elapsed not working")
	}

	Start("m3")
	time.Sleep(2 * time.Millisecond)
	d3 := Elapsed("m3")
	if d3.Milliseconds() < 1 {
		t.Errorf("timing measure cumulative not working")
	}

	time.Sleep(2 * time.Millisecond)
	d4 := Elapsed("m3")
	if d4.Milliseconds() <= d3.Milliseconds() {
		t.Errorf("timing measure cumulative not working")
	}
}

func TestTimingElapsedAll(t *testing.T) {
	tm := NewTimers("test")
	tm.Start("m1", "m2")
	time.Sleep(2 * time.Millisecond)
	elapsed1 := tm.ElapsedAll()

	time.Sleep(2 * time.Millisecond)
	elapsed2 := tm.ElapsedAll()

	if elapsed1["m1"] >= elapsed2["m1"] || elapsed1["m2"] >= elapsed2["m2"] {
		t.Errorf("elapsed all not working")
	}
}

func TestMeasureAll(t *testing.T) {
	tm := GetTimers()
	tm.Start("m1")
	time.Sleep(2 * time.Millisecond)
	t1 := tm.Measure("m1")

	tm.Start("m2")
	time.Sleep(2 * time.Millisecond)
	t2 := tm.Measure("m2")

	tm.Start("m3")
	time.Sleep(2 * time.Millisecond)

	elapsed := tm.MeasureAll()
	if elapsed["m1"] != t1 || elapsed["m2"] != t2 {
		t.Errorf("measure all not working")
	}
	if elapsed["m3"] <= time.Millisecond {
		t.Errorf("measure all not working")
	}

	elapsed2 := MeasureAll()
	if elapsed2["m1"] != t1 || elapsed2["m2"] != t2 {
		t.Errorf("timing measure all not working")
	}
	if elapsed2["m3"] <= time.Millisecond {
		t.Errorf("timing measure all not working")
	}
}

func TestPauseResume(t *testing.T) {
	tm := GetTimers()

	tm.Start("m1")
	time.Sleep(2 * time.Millisecond)
	tm.Pause("m1")

	d1 := tm.Elapsed("m1")
	time.Sleep(2 * time.Millisecond)
	d2 := tm.Elapsed("m1")
	if d1 != d2 {
		t.Errorf("pause not working")
	}

	tm.Resume("m1")
	time.Sleep(2 * time.Millisecond)
	d3 := tm.Elapsed("m1")
	if d3 <= d2 {
		t.Errorf("resume not working")
	}

	Start("m2")
	time.Sleep(2 * time.Millisecond)
	tm.Pause("m2")
	d4 := tm.Elapsed("m2")
	time.Sleep(2 * time.Millisecond)
	d5 := tm.Elapsed("m2")
	if d4 != d5 {
		t.Errorf("timing pause not working")
	}

	tm.Resume("m2")
	time.Sleep(2 * time.Millisecond)
	d6 := tm.Elapsed("m2")
	if d6 <= d5 {
		t.Errorf("timing resume not working")
	}
}

func TestPauseAll(t *testing.T) {
	tm := NewTimers("TestPauseAll")
	tm.Start("m1", "m2", "m3")
	time.Sleep(2 * time.Millisecond)
	tm.PauseAll()
	time.Sleep(2 * time.Millisecond)
	elapsed := tm.MeasureAll()
	if elapsed["m1"] > 3*time.Millisecond || elapsed["m2"] > 3*time.Millisecond || elapsed["m3"] > 3*time.Millisecond {
		t.Errorf("pause all not working")
	}
}
