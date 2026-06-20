# 3959 — Check Good Integer

## Deskripsi

**Soal:** [3959. Check Good Integer](https://leetcode.com/problems/check-good-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3959: Check Good Integer
// https://leetcode.com/problems/check-good-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckGoodInteger(1000))
	fmt.Println(CheckGoodInteger(19))
}

// Time: O(log n)
// Space: O(1)
func CheckGoodInteger(n int) bool {
	total := 0
	for n > 0 {
		d := n % 10
		total += d * (d - 1)
		n /= 10
	}
	return total >= 50
}
```
