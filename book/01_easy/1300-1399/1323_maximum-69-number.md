# 1323 — Maximum 69 Number

## Deskripsi

**Soal:** [1323. Maximum 69 Number](https://leetcode.com/problems/maximum-69-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximum69Number(num int) int`

## Solusi Go

```go
package main

// LeetCode #1323: Maximum 69 Number
// https://leetcode.com/problems/maximum-69-number/
// Difficulty: Easy
//
// LeetCode submission: func maximum69Number(num int) int

import "fmt"

func main() {
	fmt.Println(MaximumSixNineNumber(9669)) // 9969
	fmt.Println(MaximumSixNineNumber(9996)) // 9999
	fmt.Println(MaximumSixNineNumber(9999)) // 9999
}

// Time: O(log n), Space: O(1)
func MaximumSixNineNumber(num int) int {
	maxBase := 0
	base := 1
	x := num
	for x > 0 {
		if x%10 == 6 {
			maxBase = base
		}
		x /= 10
		base *= 10
	}
	return num + maxBase*3
}
```
