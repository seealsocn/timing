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
    timing.Start("total")
    timing.Start("task13")

    timing.Start("task1")
    time.Sleep(time.Millisecond)
    task1Elapsed := timing.Measure("task1")

    timing.Pause("task13")

    timing.Start("task2")
    time.Sleep(time.Millisecond)
    task2Elapsed := timing.Measure("task2")

    timing.Resume("task13")

    timing.Start("task3")
    time.Sleep(time.Millisecond)
    task3Elapsed := timing.Measure("task3")

    task13Elapsed := timing.Measure("task13")
    totalElapsed := timing.Measure("total")

    fmt.Printf("Task elapsed\ntask1: %v\ntask2: %v\ntask3: %v\ntask13: %v\ntotal: %v",
        task1Elapsed, task2Elapsed, task3Elapsed, task13Elapsed, totalElapsed)
}
```
