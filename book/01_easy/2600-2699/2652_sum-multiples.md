# 2652 — Sum Multiples

## Deskripsi

**Soal:** [2652. Sum Multiples](https://leetcode.com/problems/sum-multiples/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2652: Sum Multiples
// https://leetcode.com/problems/sum-multiples/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SumMultiples(7))
	fmt.Println(SumMultiples(10))
}

func SumMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}
```
