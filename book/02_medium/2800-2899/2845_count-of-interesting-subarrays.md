# 2845 — Count Of Interesting Subarrays

## Deskripsi

**Soal:** [2845. Count Of Interesting Subarrays](https://leetcode.com/problems/count-of-interesting-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func CountOfInterestingSubarrays(nums []int, modulo int, k int) int64`

## Solusi Go

```go
package main

// LeetCode #2845: Count of Interesting Subarrays
// https://leetcode.com/problems/count-of-interesting-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountOfInterestingSubarrays(nums []int, modulo int, k int) int64 {
  // Membuat map untuk pencarian O(1): key → value
	prefix := make(map[int]int64)
	prefix[0] = 1
	var count int64
	var sum int

	for _, num := range nums {
		if num%modulo == k {
			sum++
		}
		need := (sum%modulo - k%modulo + modulo) % modulo
		count += prefix[need]
		prefix[sum%modulo]++
	}

	return count
}

func main() {
	fmt.Println(CountOfInterestingSubarrays([]int{3, 2, 4}, 2, 1))
	fmt.Println(CountOfInterestingSubarrays([]int{1, 2, 3, 4}, 3, 1))
}
```
