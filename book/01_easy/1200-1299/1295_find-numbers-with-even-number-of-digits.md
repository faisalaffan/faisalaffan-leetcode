# 1295 — Find Numbers With Even Number Of Digits

## Deskripsi

**Soal:** [1295. Find Numbers With Even Number Of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1295: Find Numbers with Even Number of Digits
// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896})) // 2
	fmt.Println(findNumbers([]int{555, 901, 482, 1771})) // 1
}

// LeetCode submission: findNumbers
func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		digits := 0
		for x := v; x > 0; x /= 10 {
			digits++
		}
		if digits%2 == 0 {
			count++
		}
	}
	return count
}
```
