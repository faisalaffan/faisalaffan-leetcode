# 1320 — Minimum Distance To Type A Word Using Two Fingers

## Deskripsi

**Soal:** [1320. Minimum Distance To Type A Word Using Two Fingers](https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minimumDistance(word string) int`

> **Ide Kunci:** DP tracking one finger position.

## Solusi Go

```go
package main

// LeetCode #1320: Minimum Distance to Type a Word Using Two Fingers
// https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/
// Difficulty: Hard
//
// Approach: DP tracking one finger position.
// dp[i][other]: min movement to type word[0..i] where one finger is at word[i]
// and the other is at alphabetical index `other` (26 = unused).
// At each step, either the same finger types the next character, or the other
// finger moves from its current position to type the next character.
// Keyboard layout: A=(0,0), B=(0,1), ..., Z=(4,1).

import "fmt"

func minimumDistance(word string) int {
	n := len(word)
	if n <= 1 {
		return 0
	}

	dist := func(a, b byte) int {
		if a == 0 || b == 0 {
			return 0
		}
		ax, ay := int((a-'A')/6), int((a-'A')%6)
		bx, by := int((b-'A')/6), int((b-'A')%6)
		return abs(ax-bx) + abs(ay-by)
	}

	unused := 26
	INF := 1 << 30
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 27)
		for j := 0; j <= 26; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][unused] = 0

	for i := 0; i < n-1; i++ {
		cur := word[i]
		nxt := word[i+1]
		for other := 0; other <= 26; other++ {
			if dp[i][other] >= INF {
				continue
			}
			// Same finger types nxt
			cost := dp[i][other] + dist(cur, nxt)
			if cost < dp[i+1][other] {
				dp[i+1][other] = cost
			}
			// Other finger types nxt
			cost2 := dp[i][other]
			if other != unused {
				cost2 += dist(byte('A'+other), nxt)
			}
			otherIdx := int(cur - 'A')
			if cost2 < dp[i+1][otherIdx] {
				dp[i+1][otherIdx] = cost2
			}
		}
	}

	ans := INF
	for other := 0; other <= 26; other++ {
		if dp[n-1][other] < ans {
			ans = dp[n-1][other]
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minimumDistance("CAKE"))   // 3
	fmt.Println(minimumDistance("HAPPY"))  // 6
	fmt.Println(minimumDistance("A"))      // 0
	fmt.Println(minimumDistance("NEW"))    // 3
}
```
