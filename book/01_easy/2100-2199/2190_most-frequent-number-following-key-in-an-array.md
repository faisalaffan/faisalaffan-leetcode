# 2190 — Most Frequent Number Following Key In An Array

## Deskripsi

**Soal:** [2190. Most Frequent Number Following Key In An Array](https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2190: Most Frequent Number Following Key In an Array
// https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{1, 100, 200, 1, 100}, 1))    // 100
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{2, 2, 2, 2, 3}, 2))          // 2
}

// Time: O(n), Space: O(n)
func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}

	maxCount := 0
	result := 0
	for num, count := range freq {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	return result
}
```
