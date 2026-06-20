# 1512 — Number Of Good Pairs

## Deskripsi

**Soal:** [1512. Number Of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numIdenticalPairs(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1512: Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/
// Difficulty: Easy
//
// LeetCode submission: func numIdenticalPairs(nums []int) int

import "fmt"

func main() {
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println(NumberOfGoodPairs([]int{1, 1, 1, 1}))        // 6
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3}))           // 0
}

// Time: O(n), Space: O(n)
func NumberOfGoodPairs(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v]
		freq[v]++
	}
	return count
}
```
