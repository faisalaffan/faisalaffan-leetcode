# 2237 — Count Positions On Street With Required Brightness

## Deskripsi

**Soal:** [2237. Count Positions On Street With Required Brightness](https://leetcode.com/problems/count-positions-on-street-with-required-brightness/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func meetRequirement(n int, lights [][]int, requirement []int) int`

## Solusi Go

```go
package main

// LeetCode #2237: Count Positions on Street With Required Brightness
// https://leetcode.com/problems/count-positions-on-street-with-required-brightness/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func meetRequirement(n int, lights [][]int, requirement []int) int {
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n+2)
	for _, l := range lights {
		pos, r := l[0], l[1]
		left := pos - r
		if left < 0 {
			left = 0
		}
		right := pos + r
		if right > n {
			right = n
		}
		diff[left]++
		diff[right+1]--
	}

	count := 0
	brightness := 0
	for i := 0; i < n; i++ {
		brightness += diff[i]
		if brightness >= requirement[i] {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(meetRequirement(5, [][]int{{0, 1}, {2, 1}, {3, 2}}, []int{0, 2, 1, 4, 1}))
	// Expected: 3

	// Test case 2
	fmt.Println(meetRequirement(3, [][]int{{1, 1}}, []int{0, 1, 0}))
	// Expected: 2
}
```
