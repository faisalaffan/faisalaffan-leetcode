# 2757 — Generate Circular Array Values

## Deskripsi

**Soal:** [2757. Generate Circular Array Values](https://leetcode.com/problems/generate-circular-array-values/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func GenerateCircularArrayValues(arr []int, start int, count int) []int`

## Solusi Go

```go
package main

// LeetCode #2757: Generate Circular Array Values
// https://leetcode.com/problems/generate-circular-array-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func GenerateCircularArrayValues(arr []int, start int, count int) []int {
	n := len(arr)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = arr[(start+i)%n]
	}
	return result
}

func main() {
	fmt.Println(GenerateCircularArrayValues([]int{1, 2, 3, 4}, 2, 6))
	fmt.Println(GenerateCircularArrayValues([]int{10, 20}, 1, 3))
}
```
