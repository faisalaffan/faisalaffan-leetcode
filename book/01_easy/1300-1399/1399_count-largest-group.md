# 1399 — Count Largest Group

## Deskripsi

**Soal:** [1399. Count Largest Group](https://leetcode.com/problems/count-largest-group/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func countLargestGroup(n int) int`

## Solusi Go

```go
package main

// LeetCode #1399: Count Largest Group
// https://leetcode.com/problems/count-largest-group/
// Difficulty: Easy
//
// LeetCode submission: func countLargestGroup(n int) int

import "fmt"

func main() {
	fmt.Println(CountLargestGroup(13)) // 4
	fmt.Println(CountLargestGroup(2))  // 2
	fmt.Println(CountLargestGroup(15)) // 6
}

// Time: O(n log n), Space: O(n)
func CountLargestGroup(n int) int {
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[int]int)
	maxSize := 0
	for i := 1; i <= n; i++ {
		s := digitSum(i)
		groups[s]++
		if groups[s] > maxSize {
			maxSize = groups[s]
		}
	}
	count := 0
	for _, v := range groups {
		if v == maxSize {
			count++
		}
	}
	return count
}

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
```
