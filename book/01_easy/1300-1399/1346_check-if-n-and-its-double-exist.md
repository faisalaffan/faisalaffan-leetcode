# 1346 — Check If N And Its Double Exist

## Deskripsi

**Soal:** [1346. Check If N And Its Double Exist](https://leetcode.com/problems/check-if-n-and-its-double-exist/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func checkIfExist(arr []int) bool`

## Solusi Go

```go
package main

// LeetCode #1346: Check If N and Its Double Exist
// https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Difficulty: Easy
//
// LeetCode submission: func checkIfExist(arr []int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfNAndItsDoubleExist([]int{10, 2, 5, 3}))  // true (10 = 2*5)
	fmt.Println(CheckIfNAndItsDoubleExist([]int{3, 1, 7, 11}))   // false
	fmt.Println(CheckIfNAndItsDoubleExist([]int{7, 1, 14, 11}))  // true (14 = 2*7)
}

// Time: O(n), Space: O(n)
func CheckIfNAndItsDoubleExist(arr []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		if seen[v*2] || (v%2 == 0 && seen[v/2]) {
			return true
		}
		seen[v] = true
	}
	return false
}
```
