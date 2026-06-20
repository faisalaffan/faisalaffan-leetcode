# 2778 — Sum Of Squares Of Special Elements

## Deskripsi

**Soal:** [2778. Sum Of Squares Of Special Elements](https://leetcode.com/problems/sum-of-squares-of-special-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2778: Sum of Squares of Special Elements
// https://leetcode.com/problems/sum-of-squares-of-special-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SumOfSquaresOfSpecialElements([]int{1, 2, 3, 4}))
	fmt.Println(SumOfSquaresOfSpecialElements([]int{2, 7, 1, 19, 18, 3}))
}

func SumOfSquaresOfSpecialElements(nums []int) int {
	n := len(nums)
	sum := 0
	for i, v := range nums {
		if n%(i+1) == 0 {
			sum += v * v
		}
	}
	return sum
}
```
