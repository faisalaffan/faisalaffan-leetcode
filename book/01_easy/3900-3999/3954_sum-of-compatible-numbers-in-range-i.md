# 3954 — Sum Of Compatible Numbers In Range I

## Deskripsi

**Soal:** [3954. Sum Of Compatible Numbers In Range I](https://leetcode.com/problems/sum-of-compatible-numbers-in-range-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(k)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3954: Sum of Compatible Numbers in Range I
// https://leetcode.com/problems/sum-of-compatible-numbers-in-range-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfCompatibleNumbersInRangeI(2, 3))
	fmt.Println(SumOfCompatibleNumbersInRangeI(5, 1))
}

// Time: O(k)
// Space: O(1)
func SumOfCompatibleNumbersInRangeI(n int, k int) int {
	sum := 0
	start := n - k
	if start < 1 {
		start = 1
	}
	end := n + k
	for x := start; x <= end; x++ {
		if n&x == 0 {
			sum += x
		}
	}
	return sum
}
```
