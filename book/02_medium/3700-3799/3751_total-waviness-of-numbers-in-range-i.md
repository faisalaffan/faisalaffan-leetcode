# 3751 — Total Waviness Of Numbers In Range I

## Deskripsi

**Soal:** [3751. Total Waviness Of Numbers In Range I](https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * D)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func totalWavinessOfNumbersInRangeI(num1 int, num2 int) int`

## Solusi Go

```go
package main

// LeetCode #3751: Total Waviness of Numbers in Range I
// https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/
// Difficulty: Medium
// Time: O(N * D) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func totalWavinessOfNumbersInRangeI(num1 int, num2 int) int {
	ans := 0
	for num := num1; num <= num2; num++ {
		s := strconv.Itoa(num)
		for i := 1; i < len(s)-1; i++ {
			if (s[i] > s[i-1] && s[i] > s[i+1]) || (s[i] < s[i-1] && s[i] < s[i+1]) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(totalWavinessOfNumbersInRangeI(1, 100))
	fmt.Println(totalWavinessOfNumbersInRangeI(100, 200))
	fmt.Println(totalWavinessOfNumbersInRangeI(1000, 1050))
}
```
