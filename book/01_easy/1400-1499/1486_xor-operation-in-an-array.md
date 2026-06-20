# 1486 — Xor Operation In An Array

## Deskripsi

**Soal:** [1486. Xor Operation In An Array](https://leetcode.com/problems/xor-operation-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func xorOperation(n int, start int) int`

## Solusi Go

```go
package main

// LeetCode #1486: XOR Operation in an Array
// https://leetcode.com/problems/xor-operation-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func xorOperation(n int, start int) int

import "fmt"

func main() {
	fmt.Println(XorOperationInAnArray(5, 0)) // 8
	fmt.Println(XorOperationInAnArray(4, 3)) // 8
	fmt.Println(XorOperationInAnArray(1, 7)) // 7
}

// Time: O(n), Space: O(1)
func XorOperationInAnArray(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
```
