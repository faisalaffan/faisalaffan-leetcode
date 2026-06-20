# 3267 — Count Almost Equal Pairs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countAlmostEqualPairsII(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3267: Count Almost Equal Pairs II
// https://leetcode.com/problems/count-almost-equal-pairs-ii/
// Difficulty: Hard
//
// Two numbers are "almost equal" if they can be made equal by swapping
// at most one pair of digits in at most one of the numbers.
// That means either they are already equal, or they differ in exactly 2
// positions where the differing digits are swapped.

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Example 1
	fmt.Println(countAlmostEqualPairsII([]int{1, 10, 100}))
	// Example 2
	fmt.Println(countAlmostEqualPairsII([]int{3, 12, 33, 123}))
	// Example 3: all equal
	fmt.Println(countAlmostEqualPairsII([]int{1, 1, 1, 1}))
	// Example 4
	fmt.Println(countAlmostEqualPairsII([]int{123, 321, 213, 132}))
	// Example 5: single element
	fmt.Println(countAlmostEqualPairsII([]int{5}))
}

func countAlmostEqualPairsII(nums []int) int64 {
	// Group numbers by their sorted digit multiset.
	// Only numbers with the same multiset can be almost equal.
  // HashMap: O(1) lookup
	groups := make(map[string][]string)
	for _, num := range nums {
		s := strconv.Itoa(num)
		b := []byte(s)
  // Custom sort
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], s)
	}

	var ans int64

	for _, group := range groups {
		// Count frequency of each distinct string in this group.
  // HashMap: O(1) lookup
		freq := make(map[string]int)
		for _, s := range group {
			freq[s]++
		}

		// Equal pairs: any two identical numbers.
		for _, f := range freq {
			ans += int64(f) * int64(f-1) / 2
		}

		// Almost-equal (one-swap) pairs among different numbers.
		// For each number, generate all variants reachable by one swap.
		// If a variant exists in freq and is not the original number,
		// then this number and that variant form a valid pair.
		for s, f := range freq {
			b := []byte(s)
			n := len(b)
  // HashMap: O(1) lookup
			seen := make(map[string]bool)
			for p := 0; p < n; p++ {
				for q := p + 1; q < n; q++ {
					b[p], b[q] = b[q], b[p]
					variant := string(b)
					if !seen[variant] && freq[variant] > 0 && variant != s {
						ans += int64(f) * int64(freq[variant])
						seen[variant] = true
					}
					b[p], b[q] = b[q], b[p] // restore
				}
			}
		}
	}

	// Every one-swap pair was counted twice (a→b and b→a).
	ans /= 2
	return ans
}
```
