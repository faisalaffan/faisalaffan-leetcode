# 1900 — The Earliest And Latest Rounds Where Players Compete

## Deskripsi

**Soal:** [1900. The Earliest And Latest Rounds Where Players Compete](https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func dfs(a, b, n int) []int`

## Solusi Go

```go
package main

// LeetCode #1900: The Earliest and Latest Rounds Where Players Compete
// https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/
// Difficulty: Hard

import "fmt"

type state struct {
	a, b, n int
}

var memo map[state][]int

func dfs(a, b, n int) []int {
	// Returns [earliest, latest] round (1 for "this round") where a and b meet.
	if a+b == n-1 {
		return []int{1, 1}
	}

	key := state{a, b, n}
	if res, ok := memo[key]; ok {
		return res
	}

	pairs := n / 2               // number of pairs
	half := (n + 1) / 2          // number of players advancing
	mid := n / 2                 // middle player (only meaningful when n odd)

	pairA := min(a, n-1-a)
	pairB := min(b, n-1-b)

	// Other pair indices (0..pairs-1), excluding pairA and pairB
  // Membuat slice untuk menyimpan hasil
	otherPairs := make([]int, 0)
	for i := 0; i < pairs; i++ {
		if i != pairA && i != pairB {
			otherPairs = append(otherPairs, i)
		}
	}

	minRound := 100
	maxRound := 0
	m := len(otherPairs)

	for mask := 0; mask < (1 << m); mask++ {
  // Membuat slice untuk menyimpan hasil
		winners := make([]int, half)
  // Iterasi seluruh elemen
		for i := range winners {
			winners[i] = -1
		}

		// a and b always win their matches
		winners[pairA] = a
		winners[pairB] = b

		// For each other pair, try both winners
		for k, pi := range otherPairs {
			if mask&(1<<k) != 0 {
				winners[pi] = n - 1 - pi // higher-indexed player wins
			} else {
				winners[pi] = pi // lower-indexed player wins
			}
		}

		// When n is odd, the middle player (position = pairs = n/2) gets a bye
		// and advances to the last position in the new lineup.
		if n%2 == 1 {
			winners[half-1] = mid
		}

		newA, newB := -1, -1
		for i, w := range winners {
			if w == a {
				newA = i
			}
			if w == b {
				newB = i
			}
		}
		if newA != -1 && newB != -1 {
			if newA > newB {
				newA, newB = newB, newA
			}
			sub := dfs(newA, newB, half)
			earliest := 1 + sub[0]
			latest := 1 + sub[1]
			if earliest < minRound {
				minRound = earliest
			}
			if latest > maxRound {
				maxRound = latest
			}
		}
	}

	result := []int{minRound, maxRound}
	memo[key] = result
	return result
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	memo = make(map[state][]int)
	a, b := firstPlayer-1, secondPlayer-1
	if a > b {
		a, b = b, a
	}
	return dfs(a, b, n)
}

func main() {
	// Example: n=11, firstPlayer=2, secondPlayer=4 -> [3,4]
	fmt.Println(earliestAndLatest(11, 2, 4))

	// Additional test
	fmt.Println(earliestAndLatest(5, 1, 5))
}
```
