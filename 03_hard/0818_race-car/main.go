package main

// LeetCode #818: Race Car
// https://leetcode.com/problems/race-car/
// Difficulty: Hard
//
// Your car starts at position 0 with speed 1. On each step you can:
//   - 'A': position += speed; speed *= 2
//   - 'R': speed = (speed > 0 ? -1 : 1); position unchanged
// Find the shortest instruction sequence to reach target.
//
// Approach: DP
//   - Let dp[t] be the minimum instructions to reach position t.
//   - With k consecutive 'A' moves starting from (0,1), we reach position
//     2^k - 1 at speed 2^k.
//   - If 2^k - 1 == t, we're done with k instructions.
//   - If overshoot (2^k - 1 > t), reverse and in the reverse direction we must
//     cover the remaining distance: dp[t] = k + 1 + dp[(2^k - 1) - t].
//   - If undershoot (2^k - 1 < t), we can reverse early (after k-1 A's), go
//     some m A's in reverse, reverse again, and cover the rest forward.

import "fmt"

func main() {
	// target=3 → "AA" = 2
	fmt.Println(racecar(3))
	// target=6 → "AAARA" = 5 (or "AAARA")
	fmt.Println(racecar(6))
	// target=1 → "A" = 1
	fmt.Println(racecar(1))
	// target=2 → "AARA" or "ARAA" → wait, "AA" pos=3, "AR" pos=-1, "ARA" pos=0...
	// target=2: "AARA" pos: 1,3,3,2 = 4 instructions
	fmt.Println(racecar(2))
	// target=4 → "AARA" + something
	fmt.Println(racecar(4))
	// target=0 → 0
	fmt.Println(racecar(0))
}

func racecar(target int) int {
	if target == 0 {
		return 0
	}

	dp := make([]int, target+3)
	for t := 1; t <= target; t++ {
		// Find smallest k where (1<<k)-1 >= t
		k := 1
		for (1<<k)-1 < t {
			k++
		}

		exact := (1 << k) - 1
		if exact == t {
			dp[t] = k
			continue
		}

		// Option 1: Overshoot then come back
		// k A's → position exact (overshoot by exact-t), R, then come back dp[exact-t]
		dp[t] = k + 1 + dp[exact-t]

		// Option 2: Undershoot strategy
		// (k-1) A's → position (1<<(k-1))-1, R, then m A's back, R, then forward
		for m := 0; m < k-1; m++ {
			// Position after: (1<<(k-1))-1 - (1<<m) + 1 = (1<<(k-1)) - (1<<m)
			pos := (1 << (k - 1)) - (1 << m)
			if pos < t {
				remaining := t - pos
				// Instructions: (k-1) A + 1 R + m A + 1 R + dp[remaining]
				steps := (k - 1) + 1 + m + 1 + dp[remaining]
				if steps < dp[t] {
					dp[t] = steps
				}
			}
		}
	}
	return dp[target]
}
