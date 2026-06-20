# 0258 — Add Digits

## Deskripsi

**Soal:** [0258. Add Digits](https://leetcode.com/problems/add-digits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func AddDigits(num int) int`

## Solusi Go

```go
package main

// LeetCode #258: Add Digits
// https://leetcode.com/problems/add-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func AddDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	fmt.Println(AddDigits(38))
	fmt.Println(AddDigits(0))
	fmt.Println(AddDigits(9))
}
```
