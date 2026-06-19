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
	// a and b are positions (0-indexed), a < b
	// n is the current number of players
	// Returns [earliest, latest] round where a and b can meet
	if a == 0 && b == n-1 {
		// they are at the ends: they meet in the next round
		return []int{1, 1}
	}

	key := state{a, b, n}
	if res, ok := memo[key]; ok {
		return res
	}

	// Generate all possible results of a single round
	// n players, paired as (0, n-1), (1, n-2), ...
	// After the round, the winners advance.
	// The mapping from old position to new position:
	// If player at position i wins, new position = min(i, n-1-i)
	// (lower index pair, i.e., which of the floor(n/2) pairs they are in)

	pairs := n / 2
	half := (n + 1) / 2 // number of winners

	// a and b are in the same pair when a + b == n - 1
	if a+b == n-1 {
		// They compete against each other this round
		memo[key] = []int{1, 1}
		return []int{1, 1}
	}

	// Determine which pair a and b are in
	pairA := min(a, n-1-a)
	pairB := min(b, n-1-b)

	// a and b are the two players' indices in the pairing.
	// Each player can either win or lose. We need to consider all scenarios
	// where BOTH a and b win (so they advance).
	// After they win, their new positions are pairA and pairB (which are < half).

	// The other players: for each pair (i, n-1-i), exactly one advances.
	// We need to determine possible new positions for a and b after the round.

	// This is a combinatorial problem: how many players before
	// and after can survive, affecting the relative positions?

	// The full DP is complex. Let's implement the known approach:
	// The key insight: after each round, the positions of surviving players
	// are determined by their original pair indices. If a and b win,
	// they become at positions (some function of how many lower-pair players survived).

	// For two players to meet at round r:
	// earliest = shortest path to meet = minimal rounds
	// latest = longest path to meet = maximal rounds

	// Known solution: use recursion on the compressed positions.
	// The state is (a, b, n).
	// A player at position i wins -> new position = min(i, n-1-i).
	// Among the floor(n/2) pairs, exactly one per pair advances.
	// a and b win (they must both survive), so they advance.
	// We need to select for each other pair which player wins,
	// to make a and b as close together as possible (for earliest)
	// or as far apart as possible (for latest).

	// The "earliest" meeting time = 1 + min over possible outcomes of earliest in new state
	// The "latest" meeting time = 1 + max over possible outcomes of latest in new state

	// Since n <= 28, we can brute-force all possible outcomes.
	pairA = min(a, n-1-a)
	pairB = min(b, n-1-b)

	// Each of the other floor(n/2) pairs can have either player win,
	// except pairs involving a or b (where only a/b can win).
	// Wait - a and b win, so their opponents lose. But other pairs are free.

	// Actually, if a is in position a, and a wins, then the opponent at n-1-a loses.
	// So for pair pairA, only a advances.
	// Similarly for pair pairB.

	minRound := 100
	maxRound := 0

	// Generate all 2^(pairs-2) possibilities (pairs up to 14, so 2^12 = 4096 max)
	// We need to consider the relative ordering of winners.
	// The new sequence is: from each pair (i, n-1-i), the winner
	// takes the position corresponding to i (0-indexed) in the new array.
	// a and b are always winners.

	otherPairs := make([]int, 0)
	for i := 0; i < pairs; i++ {
		if i == pairA || i == pairB {
			continue
		}
		otherPairs = append(otherPairs, i)
	}

	// winners[newPos] = original player position for the winner of that pair
	// For a and b, they definitely win.
	// For other pairs, we try both possibilities.

	m := len(otherPairs)
	for mask := 0; mask < (1 << m); mask++ {
		winners := make([]int, half)
		for i := 0; i < half; i++ {
			winners[i] = -1
		}
		winners[pairA] = a
		winners[pairB] = b

		for k, i := range otherPairs {
			if mask&(1<<k) != 0 {
				// The higher-indexed player wins
				winners[i] = n - 1 - i
			} else {
				// The lower-indexed player wins
				winners[i] = i
			}
		}

		// Now find new positions of a and b in the winners array
		newA := -1
		newB := -1
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
