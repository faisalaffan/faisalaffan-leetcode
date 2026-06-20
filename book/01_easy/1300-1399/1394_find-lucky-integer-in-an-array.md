# 1394 — Find Lucky Integer In An Array

## Deskripsi

**Soal:** [1394. Find Lucky Integer In An Array](https://leetcode.com/problems/find-lucky-integer-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findLucky(arr []int) int`

## Solusi Go

```go
package main

// LeetCode #1394: Find Lucky Integer in an Array
// https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func findLucky(arr []int) int

import "fmt"

func main() {
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 3, 4}))       // 2
	fmt.Println(FindLuckyIntegerInAnArray([]int{1, 2, 2, 3, 3, 3})) // 3
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 2, 3, 3}))    // -1
}

// Time: O(n), Space: O(n)
func FindLuckyIntegerInAnArray(arr []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int, len(arr))
	for _, v := range arr {
		freq[v]++
	}
	ans := -1
	for k, v := range freq {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
```
