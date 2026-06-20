# 1437 — Check If All 1S Are At Least Length K Places Away

## Deskripsi

**Soal:** [1437. Check If All 1S Are At Least Length K Places Away](https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func kLengthApart(nums []int, k int) bool`

## Solusi Go

```go
package main

// LeetCode #1437: Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
// Difficulty: Easy
//
// LeetCode submission: func kLengthApart(nums []int, k int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2)) // true
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 1, 0, 1}, 2))       // false
}

// Time: O(n), Space: O(1)
func CheckIfAllOneSAreAtLeastLengthKPlacesAway(nums []int, k int) bool {
	prev := -k - 1
	for i, v := range nums {
		if v == 1 {
			if i-prev-1 < k {
				return false
			}
			prev = i
		}
	}
	return true
}
```
