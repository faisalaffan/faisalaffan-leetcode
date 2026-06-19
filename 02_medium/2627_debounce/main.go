package main

// LeetCode #2627: Debounce
// https://leetcode.com/problems/debounce/
// Difficulty: Medium
// Time: O(1) per call | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type DebouncedFn struct {
	mu       sync.Mutex
	timer    *time.Timer
	duration time.Duration
	fn       func(...int)
}

func (d *DebouncedFn) Call(args ...int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.duration, func() {
		d.fn(args...)
	})
}

func (d *DebouncedFn) Cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
}

func newDebounce(fn func(...int), duration time.Duration) *DebouncedFn {
	return &DebouncedFn{
		fn:       fn,
		duration: duration,
	}
}

func main() {
	calls := []int{}
	debounced := newDebounce(func(args ...int) {
		calls = append(calls, args[0])
	}, 50*time.Millisecond)

	debounced.Call(1)
	debounced.Call(2)
	debounced.Call(3)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Test 1:", calls)
	// Expected: [3] (only last call goes through)

	debounced.Call(4)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Test 2:", calls)
	// Expected: [3, 4]

	debounced.Cancel()
	fmt.Println("Test 3: cancelled without call")
}
