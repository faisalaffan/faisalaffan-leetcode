# 2449 — Minimum Number Of Operations To Make Arrays Similar

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func makeSimilar(nums []int, target []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2449: Minimum Number of Operations to Make Arrays Similar
// https://leetcode.com/problems/minimum-number-of-operations-to-make-arrays-similar/
// Difficulty: Hard
//
// Sort both arrays, separate by parity (odd/even). Each operation changes
// a value by +/-2 preserving parity. Match elements of same parity in sorted
// order. Count the total positive difference (sum of target[i] - nums[i] for
// those needing increase). Each operation fixes 2 units, so ans = totalPos / 2.
// Time O(N log N) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))
	// Example 2
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))
	// Already similar
	fmt.Println(makeSimilar([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}))
}

func makeSimilar(nums []int, target []int) int64 {
	numsOdd := collectOdd(nums)
	tgtOdd := collectOdd(target)
	numsEven := collectEven(nums)
	tgtEven := collectEven(target)

  // Sort O(n log n)
	sort.Ints(numsOdd)
  // Sort O(n log n)
	sort.Ints(tgtOdd)
  // Sort O(n log n)
	sort.Ints(numsEven)
  // Sort O(n log n)
	sort.Ints(tgtEven)

	var posDiff int64
  // Linear scan O(n)
	for i := 0; i < len(numsOdd); i++ {
		if tgtOdd[i] > numsOdd[i] {
			posDiff += int64(tgtOdd[i]-numsOdd[i]) / 2
		}
	}
  // Linear scan O(n)
	for i := 0; i < len(numsEven); i++ {
		if tgtEven[i] > numsEven[i] {
			posDiff += int64(tgtEven[i]-numsEven[i]) / 2
		}
	}
	return posDiff
}

func collectOdd(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 != 0 {
			res = append(res, v)
		}
	}
	return res
}

func collectEven(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}
```
