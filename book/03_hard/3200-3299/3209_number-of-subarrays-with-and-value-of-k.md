# 3209 — Number Of Subarrays With And Value Of K

## Deskripsi

**Soal:** [3209. Number Of Subarrays With And Value Of K](https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64`

> **Ide Kunci:** maintain map of (AND value -> count) for subarrays ending at the

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	cur := make(map[int]int)

	for _, x := range nums {
  // Membuat map untuk pencarian O(1): key → value
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
