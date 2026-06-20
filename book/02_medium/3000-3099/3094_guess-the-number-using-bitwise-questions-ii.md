# 3094 — Guess The Number Using Bitwise Questions Ii

## Deskripsi

**Soal:** [3094. Guess The Number Using Bitwise Questions Ii](https://leetcode.com/problems/guess-the-number-using-bitwise-questions-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func commonBits(num int) int`

## Solusi Go

```go
package main

// LeetCode #3094: Guess the Number Using Bitwise Questions II
// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/
// https://leetcode.com/problems/guess-the-number-using-bitwise-questions-ii/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

var hiddenNumber int

func commonBits(num int) int {
	count := 0
	x := hiddenNumber ^ num
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return 30 - count
}

func findNumber() int {
	n := 0
	for i := 0; i <= 30; i++ {
		bit1 := commonBits(n | (1 << i))
		bit0 := commonBits(n)
		if bit1 > bit0 {
			n |= (1 << i)
		}
	}
	return n
}

func main() {
	hiddenNumber = 42
	fmt.Println(findNumber())

	hiddenNumber = 100
	fmt.Println(findNumber())
}
```
