# 3726 — Remove Zeros In Decimal Representation

## Deskripsi

**Soal:** [3726. Remove Zeros In Decimal Representation](https://leetcode.com/problems/remove-zeros-in-decimal-representation/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n) - number of digits  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3726: Remove Zeros in Decimal Representation
// https://leetcode.com/problems/remove-zeros-in-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveZerosInDecimalRepresentation(1020030))
	fmt.Println(RemoveZerosInDecimalRepresentation(1000))
}

// Time: O(log n) - number of digits
// Space: O(1)
func RemoveZerosInDecimalRepresentation(n int) int {
	ans := 0
	k := 1
	for n > 0 {
		x := n % 10
		if x > 0 {
			ans = k*x + ans
			k *= 10
		}
		n /= 10
	}
	return ans
}
```
