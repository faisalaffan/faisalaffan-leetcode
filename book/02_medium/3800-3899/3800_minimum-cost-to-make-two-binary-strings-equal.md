# 3800 — Minimum Cost To Make Two Binary Strings Equal

## Deskripsi

**Soal:** [3800. Minimum Cost To Make Two Binary Strings Equal](https://leetcode.com/problems/minimum-cost-to-make-two-binary-strings-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func min(a, b int) int`

> **Ide Kunci:** Count mismatches and evaluate three strategies: all flips,

## Solusi Go

```go
package main

// LeetCode #3800: Minimum Cost to Make Two Binary Strings Equal
// https://leetcode.com/problems/minimum-cost-to-make-two-binary-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Approach: Count mismatches and evaluate three strategies: all flips,
// swaps+flips, or cross-swaps+swaps+flips. Take the minimum.

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinimumCostToMakeTwoBinaryStringsEqual(s string, t string, flipCost int, swapCost int, crossCost int) int {
	n := len(s)
	diff := []int{0, 0} // diff[0] = mismatches where s[i]=='0', diff[1] = mismatches where s[i]=='1'

	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			diff[int(s[i]-'0')]++
		}
	}

	totalDiff := diff[0] + diff[1]
	if totalDiff == 0 {
		return 0
	}

	ans := totalDiff * flipCost // Strategy A: flip all

	// Strategy B: pair mismatches within each s-value, swap them (costs swapCost per pair), flip remainder
	mx, mn := diff[0], diff[1]
	if mx < mn {
		mx, mn = mn, mx
	}
	ans = min(ans, mn*swapCost+(mx-mn)*flipCost)

	// Strategy C: cross-swaps (swap s[i] and t[i] at mismatched positions) + flips
	// Each cross-swap fixes both bits at a mismatched position (s[i] and t[i] trade values)
	crossPairs := totalDiff / 2
	remaining := totalDiff % 2
	ans = min(ans, crossPairs*crossCost+remaining*flipCost)

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("01000", "10111", 10, 2, 2)) // Expected: 16

	// Example 2
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("001", "110", 2, 100, 100)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToMakeTwoBinaryStringsEqual("1010", "1010", 5, 5, 5)) // Expected: 0
}
```
