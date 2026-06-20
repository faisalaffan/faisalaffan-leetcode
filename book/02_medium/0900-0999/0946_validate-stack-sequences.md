# 0946 — Validate Stack Sequences

## Deskripsi

**Soal:** [0946. Validate Stack Sequences](https://leetcode.com/problems/validate-stack-sequences/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func validateStackSequences(pushed []int, popped []int) bool`

## Solusi Go

```go
package main

// LeetCode #946: Validate Stack Sequences
// https://leetcode.com/problems/validate-stack-sequences/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func validateStackSequences(pushed []int, popped []int) bool {
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
	j := 0
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
```
