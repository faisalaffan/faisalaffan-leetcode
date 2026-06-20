# 3754 — Concatenate Non Zero Digits And Multiply By Sum I

## Deskripsi

**Soal:** [3754. Concatenate Non Zero Digits And Multiply By Sum I](https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3754: Concatenate Non-Zero Digits and Multiply by Sum I
// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenateNonZeroDigitsAndMultiplyBySumI(10203004))
	fmt.Println(ConcatenateNonZeroDigitsAndMultiplyBySumI(1000))
}

// Time: O(log n)
// Space: O(1)
func ConcatenateNonZeroDigitsAndMultiplyBySumI(n int) int {
	concat := 0
	digitSum := 0
	multiplier := 1

	for n > 0 {
		d := n % 10
		if d != 0 {
			concat = d*multiplier + concat
			multiplier *= 10
			digitSum += d
		}
		n /= 10
	}

	return concat * digitSum
}
```
