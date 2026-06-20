# 2729 — Check If The Number Is Fascinating

## Deskripsi

**Soal:** [2729. Check If The Number Is Fascinating](https://leetcode.com/problems/check-if-the-number-is-fascinating/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2729: Check if The Number is Fascinating
// https://leetcode.com/problems/check-if-the-number-is-fascinating/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(CheckIfTheNumberIsFascinating(192))
	fmt.Println(CheckIfTheNumberIsFascinating(100))
}

func CheckIfTheNumberIsFascinating(n int) bool {
	concat := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(concat) != 9 {
		return false
	}

	digits := []byte(concat)
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	return string(digits) == "123456789"
}
```
