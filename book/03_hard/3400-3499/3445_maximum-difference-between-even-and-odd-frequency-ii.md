# 3445 — Maximum Difference Between Even And Odd Frequency Ii

## Deskripsi

**Soal:** [3445. Maximum Difference Between Even And Odd Frequency Ii](https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

> **Ide Kunci:** Enumerate all pairs (a,b). Use sliding window with prefix state

## Solusi Go

```go
package main

// LeetCode #3445: Maximum Difference Between Even and Odd Frequency II
// https://leetcode.com/problems/maximum-distance-between-even-and-odd-frequency-ii/
// Difficulty: Hard
//
// Given string s (digits '0'-'4') and k, maximize freq[a] - freq[b] across
// any substring of length >= k, where freq[a] is odd and freq[b] is even.
//
// Approach: Enumerate all pairs (a,b). Use sliding window with prefix state
// compression tracking parity of counts.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(maxDifference("12223", 2))
	// Example 2
	fmt.Println(maxDifference("111", 2))
	// Example 3
	fmt.Println(maxDifference("000", 1))
	// Edge: length exactly k
	fmt.Println(maxDifference("01", 2))
	// Single digit type
	fmt.Println(maxDifference("1111", 3))
}

const INF = math.MaxInt32
const NEG_INF = math.MinInt32

func maxDifference(s string, k int) int {
	n := len(s)
	ans := NEG_INF

	for a := 0; a <= 4; a++ {
		for b := 0; b <= 4; b++ {
			if a == b {
				continue
			}
			// best[state] = min(prev_a - prev_b) for that parity state
			best := [4]int{INF, INF, INF, INF}
			cntA, cntB := 0, 0
			prevA, prevB := 0, 0
			left := -1

			for right := 0; right < n; right++ {
				dig := int(s[right] - '0')
				if dig == a {
					cntA++
				}
				if dig == b {
					cntB++
				}

				// Shrink window to maintain length >= k and cnt_b >= 2
				for right-left >= k && cntB-prevB >= 2 {
					state := ((prevA & 1) << 1) | (prevB & 1)
					val := prevA - prevB
					if val < best[state] {
						best[state] = val
					}
					left++
					if left < n && int(s[left]-'0') == a {
						prevA++
					}
					if left < n && int(s[left]-'0') == b {
						prevB++
					}
				}

				// Check current window: need cntA odd, cntB even non-zero
				if cntB >= 2 && cntA > 0 {
					rState := ((cntA & 1) << 1) | (cntB & 1)
					needState := rState ^ 2 // flip bit 1 (a parity)
					if best[needState] != INF {
						diff := (cntA - cntB) - best[needState]
						if diff > ans {
							ans = diff
						}
					}
				}
			}
		}
	}

	if ans == NEG_INF {
		return 0
	}
	return ans
}
```
