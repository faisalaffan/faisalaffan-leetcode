package main

import (
	"fmt"
	"sync"
)

// LeetCode #1226: The Dining Philosophers
// https://leetcode.com/problems/the-dining-philosophers/
// Difficulty: Medium

// 5 philosophers, 5 forks. Each needs 2 forks to eat.
// Deadlock avoidance: odd philosophers pick left then right,
// even philosophers pick right then left.

// Time: O(1) per eat call
// Space: O(1)

type DiningPhilosophers struct {
	mu     sync.Mutex
	forks  [5]sync.Mutex
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{}
}

func (dp *DiningPhilosophers) WantsToEat(philosopher int,
	eat func(),
	pickLeftFork func(),
	pickRightFork func(),
	putLeftFork func(),
	putRightFork func()) {

	left := philosopher
	right := (philosopher + 1) % 5

	// To avoid deadlock: always pick lower-numbered fork first
	if left < right {
		dp.forks[left].Lock()
		dp.forks[right].Lock()
	} else {
		dp.forks[right].Lock()
		dp.forks[left].Lock()
	}

	pickLeftFork()
	pickRightFork()
	eat()
	putLeftFork()
	putRightFork()

	dp.forks[left].Unlock()
	dp.forks[right].Unlock()
}

func main() {
	dp := NewDiningPhilosophers()
	var wg sync.WaitGroup

	eatCount := 0
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				dp.WantsToEat(p,
					func() {
						mu.Lock()
						eatCount++
						mu.Unlock()
					},
					func() { fmt.Printf("P%d picks left\n", p) },
					func() { fmt.Printf("P%d picks right\n", p) },
					func() { fmt.Printf("P%d puts left\n", p) },
					func() { fmt.Printf("P%d puts right\n", p) },
				)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Total eats: %d (expected: 10)\n", eatCount)
}
