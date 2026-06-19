package main

// LeetCode #2721: Execute Asynchronous Functions in Parallel
// https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int {
	var wg sync.WaitGroup
	results := make([]int, len(functions))

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	})
	fmt.Println(results)

	results2 := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 42 },
	})
	fmt.Println(results2)
}
