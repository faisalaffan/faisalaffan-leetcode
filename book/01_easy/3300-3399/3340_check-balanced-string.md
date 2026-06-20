# 3340 — Check Balanced String

## Deskripsi

**Soal:** [3340. Check Balanced String](https://leetcode.com/problems/check-balanced-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3340: Check Balanced String
// https://leetcode.com/problems/check-balanced-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckBalancedString("1234"))
	fmt.Println(CheckBalancedString("24123"))
}

// CheckBalancedString returns true if sum of digits at even positions equals sum at odd positions.
// Time: O(n). Space: O(1).
func CheckBalancedString(num string) bool {
	evenSum, oddSum := 0, 0
	for i, ch := range num {
		digit := int(ch - '0')
		if i%2 == 0 {
			evenSum += digit
		} else {
			oddSum += digit
		}
	}
	return evenSum == oddSum
}
```
