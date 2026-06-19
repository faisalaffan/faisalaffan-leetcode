package main

// LeetCode #2795: Parallel Execution of Promises for Individual Results Retrieval
// https://leetcode.com/problems/parallel-execution-of-promises-for-individual-results-retrieval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

type PromiseResult struct {
	Index int
	Value interface{}
}

func ParallelExecutionOfPromisesForIndividualResultsRetrieval(functions []func() int) []int {
	n := len(functions)
	results := make([]int, n)
	var wg sync.WaitGroup

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
	results := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 1 + 1 },
		func() int { return 2 + 2 },
		func() int { return 3 + 3 },
	})
	fmt.Println(results)

	results2 := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 99 },
	})
	fmt.Println(results2)
}
