# 2808 — Minimum Seconds To Equalize A Circular Array

## Deskripsi

**Soal:** [2808. Minimum Seconds To Equalize A Circular Array](https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumSecondsToEqualizeACircularArray(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2808: Minimum Seconds to Equalize a Circular Array
// https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumSecondsToEqualizeACircularArray(nums []int) int {
	n := len(nums)
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := n / 2
	for _, positions := range pos {
		if len(positions) == 0 {
			continue
		}
		maxGap := 0
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(positions); i++ {
			curr := positions[i]
			var prev int
			if i > 0 {
				prev = positions[i-1]
			} else {
				prev = positions[len(positions)-1] - n
			}
			gap := curr - prev
			if gap > maxGap {
				maxGap = gap
			}
		}
		seconds := maxGap / 2
		if seconds < best {
			best = seconds
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{1, 2, 1, 2}))
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{2, 1, 3, 3, 2}))
}
```
