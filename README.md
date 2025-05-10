# timing

A simple Go package for timing your code. Its purpose is to offer a straightforward and lightweight library that enables you to benchmark particular sections of your code whenever it's necessary.

## Example

```go
package main

import (
    "fmt"
    "time"

    "github.com/seealsocn/timing"
)

func main() {
    timers := timing.NewTimers()

    timers.Start("total")
    timers.Start("task13")

    timers.Start("task1")
    time.Sleep(time.Millisecond)
    task1Elapsed := timers.Measure("task1")

    timers.Pause("task13")

    timers.Start("task2")
    time.Sleep(time.Millisecond)
    task2Elapsed := timers.Measure("task2")

    timers.Resume("task13")

    timers.Start("task3")
    time.Sleep(time.Millisecond)
    task3Elapsed := timers.Measure("task3")

    task13Elapsed := timers.Measure("task13")
    totalElapsed := timers.Measure("total")

    fmt.Printf("Task elapsed\ntask1: %v\ntask2: %v\ntask3: %v\ntask13: %v\ntotal: %v",
        task1Elapsed, task2Elapsed, task3Elapsed, task13Elapsed, totalElapsed)

    fmt.Printf("Measure all: %+v", timers.MeasureAll())
}
```
