# 1815 — Maximum Number Of Groups Getting Fresh Donuts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxHappyGroups(batchSize int, groups []int) int
```

> **💡 Hint:** DP with Memoization (state compression via int64).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1815: Maximum Number of Groups Getting Fresh Donuts
// https://leetcode.com/problems/maximum-number-of-groups-getting-fresh-donuts/
// Difficulty: Hard
//
// Approach: DP with Memoization (state compression via int64).
//   We can reorder groups arbitrarily. A group gets fresh donuts iff the
//   running total before serving them is divisible by batchSize.
//   State = (counts of each remainder modulo batchSize, current running remainder).
//   Encode the state into an int64 and use memoized DFS.
//   Remainder-0 groups always get fresh donuts (they reset the batch),
//   so we serve them first, then recurse on non-zero remainders.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", maxHappyGroups(3, []int{1, 2, 3, 4, 5, 6}))
	// Expected: 4

	// Example 2
	fmt.Println("Example 2:", maxHappyGroups(4, []int{1, 3, 2, 5, 2, 2, 1, 6}))
	// Expected: 4

	// Edge case
	fmt.Println("Edge (single group):", maxHappyGroups(5, []int{5}))
	// Expected: 1

	// All divisible
	fmt.Println("Edge (all divisible):", maxHappyGroups(3, []int{3, 6, 9}))
	// Expected: 3
}

func maxHappyGroups(batchSize int, groups []int) int {
  // Alokasi slice integer
	counts := make([]int, batchSize)
	zeroRem := 0
	for _, g := range groups {
		r := g % batchSize
		if r == 0 {
			zeroRem++
		} else {
			counts[r]++
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[int64]int)

	var dfs func(state int64, left int) int
	dfs = func(state int64, left int) int {
		if left == 0 {
			return 0
		}
		if val, ok := memo[state]; ok {
			return val
		}
		cur := int(state & 0xF)       // current running remainder (4 bits)
		best := 0
		for r := 1; r < batchSize; r++ {
			c := int((state >> (4 + 5*(r-1))) & 0x1F)
			if c == 0 {
				continue
			}
			fresh := 0
			if cur == 0 {
				fresh = 1
			}
			newCur := (cur + r) % batchSize
			newState := state
			// decrement count for remainder r
			newState -= 1 << (4 + 5*(r-1))
			// update current remainder
			newState &^= 0xF
			newState |= int64(newCur)

			val := fresh + dfs(newState, left-1)
			if val > best {
				best = val
			}
		}
		memo[state] = best
		return best
	}

	// Build initial state: current remainder = 0, encode counts
	state := int64(0) // cur=0
	for r := 1; r < batchSize; r++ {
		state |= int64(counts[r]) << (4 + 5*(r-1))
	}

	totalLeft := 0
	for r := 1; r < batchSize; r++ {
		totalLeft += counts[r]
	}

	return zeroRem + dfs(state, totalLeft)
}

// Stub kept for compatibility with the repo scaffold.
func MaximumNumberOfGroupsGettingFreshDonuts() any {
	return maxHappyGroups(3, []int{1, 2, 3, 4, 5, 6})
}
```
