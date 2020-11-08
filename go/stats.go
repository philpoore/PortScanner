package main

import (
	"fmt"
	"sync"
)

type statsCounter struct {
	total  int
	open   int
	closed int
	mu     sync.Mutex
}

func (stats *statsCounter) update(res bool) {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	stats.total++
	if res {
		stats.open++
	} else {
		stats.closed++
	}
}

func (stats *statsCounter) display() {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	fmt.Printf("Results: total=%d closed=%d open=%d\n", stats.total, stats.closed, stats.open)
}
