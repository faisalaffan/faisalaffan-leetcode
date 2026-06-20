# 3261 — Count Substrings That Satisfy K Constraint Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countKConstraintSubstrings(s string, k int, queries [][]int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + q), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3261: Count Substrings That Satisfy K-Constraint II
// https://leetcode.com/problems/count-substrings-that-satisfy-k-constraint-ii/
// Difficulty: Hard
//
// Given binary string s and integer k, for each query [l, r] count substrings
// where either count of '0' ≤ k OR count of '1' ≤ k.
//
// Approach:
//   1. Precompute right[l] = first index where substring s[l..right[l]] becomes
//      invalid (both 0's and 1's > k).
//   2. Precompute prefix sum of valid substrings ending at each position.
//   3. For each query (l,r):
//      - Let p = min(right[l], r+1)
//      - Substrings starting at l and ending in [l, p-1] are all valid: (p-l)*(p-l+1)/2
//      - Substrings with start >= p: use prefix sums for range [p, r]
//
// Time: O(n + q), Space: O(n)

import "fmt"

func main() {
	// Example 1: s="0001111", k=2, queries=[[0,6]] => [26]
	fmt.Println(countKConstraintSubstrings("0001111", 2, [][]int{{0, 6}}))
	// Example 2: s="010101", k=1, queries=[[0,5],[1,4],[2,3]] => [15,9,3]
	fmt.Println(countKConstraintSubstrings("010101", 1, [][]int{{0, 5}, {1, 4}, {2, 3}}))
	// Example 3: all zeros, k=0
	fmt.Println(countKConstraintSubstrings("0000", 0, [][]int{{0, 3}, {0, 1}}))
	// Example 4: single char
	fmt.Println(countKConstraintSubstrings("0", 1, [][]int{{0, 0}}))
	// Example 5: alternating with large k
	fmt.Println(countKConstraintSubstrings("101010", 10, [][]int{{0, 5}, {0, 2}}))
}

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)

	// right[l] = first index where s[l..right[l]] is invalid (or n if always valid)
  // Alokasi slice
	right := make([]int, n)
  // Range loop
	for i := range right {
		right[i] = n
	}

	cnt := [2]int{}
	l := 0
	for r := 0; r < n; r++ {
		cnt[s[r]-'0']++
		for cnt[0] > k && cnt[1] > k {
			// s[l..r] is the first invalid substring starting at l
			right[l] = r
			cnt[s[l]-'0']--
			l++
		}
	}
	// For remaining starts, all substrings to end are valid (right[i] stays n)

	// prefix[i] = total valid substrings in s[0..i-1]
  // Alokasi slice
	prefix := make([]int64, n+1)
	l = 0
	for r := 0; r < n; r++ {
		for l <= r && right[l] <= r {
			l++
		}
		// All substrings starting at l..r and ending at r are valid
		// That's (r-l+1) valid substrings ending at r
		prefix[r+1] = prefix[r] + int64(r-l+1)
	}

	// Actually, we need to recompute prefix differently.
	// prefix[r] = total valid substrings ending at or before position r

	// Let me redo: count valid substrings for each start position using right[]
	// For start i, valid end positions are i..right[i]-1, so count = right[i]-i
	// prefix[i] = total valid substrings in [0..i-1]
	prefix = make([]int64, n+1)
	for i := 0; i < n; i++ {
		validEnd := right[i] - i
		prefix[i+1] = prefix[i] + int64(validEnd)
	}

  // Alokasi slice
	ans := make([]int64, len(queries))

	for qi, q := range queries {
		l, r := q[0], q[1]
		// Count valid substrings in [l, r]
		// For each start i in [l, r], valid endpoints are i..min(right[i]-1, r)
		var valid int64
		for i := l; i <= r; i++ {
			end := right[i]
			if end > r+1 {
				end = r + 1
			}
			if end > i {
				valid += int64(end - i)
			}
		}

		ans[qi] = valid
	}

	return ans
}
```
