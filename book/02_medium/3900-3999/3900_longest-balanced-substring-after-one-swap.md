# 3900 — Longest Balanced Substring After One Swap

## Deskripsi

**Soal:** [3900. Longest Balanced Substring After One Swap](https://leetcode.com/problems/longest-balanced-substring-after-one-swap/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func LongestBalancedSubstringAfterOneSwap(s string) int`

> **Ide Kunci:** Prefix sum (0->-1, 1->+1). Track first occurrence of each prefix sum.

## Solusi Go

```go
package main

// LeetCode #3900: Longest Balanced Substring After One Swap
// https://leetcode.com/problems/longest-balanced-substring-after-one-swap/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Prefix sum (0->-1, 1->+1). Track first occurrence of each prefix sum.
// Check offsets 0 (no swap) and +/-2 (one swap possible).

import "fmt"

func LongestBalancedSubstringAfterOneSwap(s string) int {
	n := len(s)

  // Membuat slice untuk menyimpan hasil
	pref := make([]int, n+1)
  // Membuat slice untuk menyimpan hasil
	onesPref := make([]int, n+1) // prefix count of '1's
	for i := 0; i < n; i++ {
		onesPref[i+1] = onesPref[i]
		if s[i] == '1' {
			pref[i+1] = pref[i] + 1
			onesPref[i+1]++
		} else {
			pref[i+1] = pref[i] - 1
		}
	}

	totalOnes := onesPref[n]
	totalZeros := n - totalOnes

  // Membuat map untuk pencarian O(1): key → value
	firstPos := make(map[int]int)
	firstPos[0] = 0

	ans := 0
	for i := 1; i <= n; i++ {
		// No swap needed (offset 0)
		if pos, ok := firstPos[pref[i]]; ok {
			if i-pos > ans {
				ans = i - pos
			}
		} else {
			firstPos[pref[i]] = i
		}

		// Swap '1' in with '0' out (offset +2, sum too many 1s)
		if pos, ok := firstPos[pref[i]-2]; ok {
			onesInside := onesPref[i] - onesPref[pos]
			zerosInside := (i - pos) - onesInside
			zerosOutside := totalZeros - zerosInside
			if onesInside > 0 && zerosOutside > 0 {
				if i-pos > ans {
					ans = i - pos
				}
			}
		}

		// Swap '0' in with '1' out (offset -2, sum too many 0s)
		if pos, ok := firstPos[pref[i]+2]; ok {
			onesInside := onesPref[i] - onesPref[pos]
			zerosInside := (i - pos) - onesInside
			onesOutside := totalOnes - onesInside
			if zerosInside > 0 && onesOutside > 0 {
				if i-pos > ans {
					ans = i - pos
				}
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestBalancedSubstringAfterOneSwap("100001")) // Expected: 4

	// Example 2
	fmt.Println(LongestBalancedSubstringAfterOneSwap("111")) // Expected: 0

	// Extra
	fmt.Println(LongestBalancedSubstringAfterOneSwap("01"))   // Expected: 2
	fmt.Println(LongestBalancedSubstringAfterOneSwap("1100")) // Expected: 4
}
```
