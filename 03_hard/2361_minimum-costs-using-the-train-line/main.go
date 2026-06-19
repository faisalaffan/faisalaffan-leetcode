package main

import (
	"fmt"
	"math"
)

// 2361. Minimum Costs Using the Train Line
// ----------------------------------------------------------------
// Two train lines: regular and express.
//   - regular[i] = cost to travel from station i to i+1 on regular.
//   - express[i] = cost to travel from station i to i+1 on express.
//   - expressCost = one‑time switch fee to move from regular to express.
//   - regularCost (same as expressCost? The problem uses two fees:
//     expressStart for regular→express, expressEnd for express→regular? Actually
//     the problem uses "expressCost" for both directions or has two separate
//     values?)
//
// Actually reading LeetCode 2361: you start on regular.  There are two
// switch costs: expressStart (regular → express) and expressEnd
// (express → regular).
//
// dpReg[i] = min cost to reach station i on the regular line.
// dpExp[i] = min cost to reach station i on the express line.
//
// Transition from station i to i+1:
//   Stay on regular:          dpReg[i] + regular[i]
//   Switch to regular:        dpExp[i] + expressEnd + regular[i]
//   Stay on express:          dpExp[i] + express[i]
//   Switch to express:        dpReg[i] + expressStart + express[i]
//
// Initialize:
//   dpReg[0] = 0
//   dpExp[0] = INF (cannot start on express)
//
// Answer: min(dpReg[n], dpExp[n]).

func minimumCosts(regular, express []int, expressStart int) []int64 {
	n := len(regular) // number of segments (from 0 to n)
	// dp for reaching station i (0-indexed)
	dpReg := int64(0)
	dpExp := int64(math.MaxInt64)

	ans := make([]int64, n)
	for i := 0; i < n; i++ {
		// Compute dp for station i+1.
		newReg := int64(math.MaxInt64)
		newExp := int64(math.MaxInt64)

		// To stay on regular or switch from express.
		if dpReg != math.MaxInt64 {
			newReg = min64(newReg, dpReg+int64(regular[i]))
		}
		if dpExp != math.MaxInt64 {
			newReg = min64(newReg, dpExp+int64(expressStart)+int64(regular[i])) // switch to regular
		}

		// To stay on express or switch from regular.
		if dpExp != math.MaxInt64 {
			newExp = min64(newExp, dpExp+int64(express[i]))
		}
		if dpReg != math.MaxInt64 {
			newExp = min64(newExp, dpReg+int64(expressStart)+int64(express[i])) // switch to express
		}

		dpReg = newReg
		dpExp = newExp

		ans[i] = min64(dpReg, dpExp)
	}
	return ans
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumCostsUsingTheTrainLine() interface{} {
	return minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
}

func main() {
	fmt.Println(MinimumCostsUsingTheTrainLine())

	got := minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8)
	want := []int64{1, 7, 14, 19}
	if len(got) != len(want) {
		fmt.Printf("FAIL: len got %d, want %d\n", len(got), len(want))
	} else {
		for i := range got {
			if got[i] != want[i] {
				fmt.Printf("FAIL [%d]: got %d, want %d\n", i, got[i], want[i])
			}
		}
	}

	got2 := minimumCosts([]int{11, 5, 13}, []int{7, 10, 6}, 3)
	want2 := []int64{10, 15, 24}
	if len(got2) != len(want2) {
		fmt.Printf("FAIL2: len got %d, want %d\n", len(got2), len(want2))
	} else {
		for i := range got2 {
			if got2[i] != want2[i] {
				fmt.Printf("FAIL2 [%d]: got %d, want %d\n", i, got2[i], want2[i])
			}
		}
	}
	fmt.Println("Done testing 2361.")
}
