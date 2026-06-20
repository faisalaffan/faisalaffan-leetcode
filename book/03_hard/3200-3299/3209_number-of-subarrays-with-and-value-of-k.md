# 3209 — Number Of Subarrays With And Value Of K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Monotonic Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3209: Number of Subarrays With AND Value of K
// https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/
// Difficulty: Hard
//
// Count subarrays whose bitwise AND equals exactly k.
// AND monotonically decreases as subarrays extend. For each ending position,
// there are at most O(log MAX) distinct AND values.
//
// Approach: maintain map of (AND value -> count) for subarrays ending at the
// current position.

import "fmt"

func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64 {
	var ans int64 = 0
  // HashMap: O(1) lookup
	cur := make(map[int]int)

	for _, x := range nums {
  // HashMap: O(1) lookup
		nxt := make(map[int]int)
		nxt[x] = 1
		for val, cnt := range cur {
			key := val & x
			nxt[key] += cnt
		}
		if cnt, ok := nxt[k]; ok {
			ans += int64(cnt)
		}
		cur = nxt
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubarraysWithAndValueOfK([]int{1, 1, 1}, 1)) // expect 6
	fmt.Println(numberOfSubarraysWithAndValueOfK([]int{1, 1, 2}, 1)) // expect ?
}
```
