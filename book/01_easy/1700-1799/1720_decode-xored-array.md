# 1720 — Decode Xored Array

## Deskripsi

**Soal:** [1720. Decode Xored Array](https://leetcode.com/problems/decode-xored-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Decode(encoded []int, first int) []int`

## Solusi Go

```go
package main

// LeetCode #1720: Decode XORed Array
// https://leetcode.com/problems/decode-xored-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func Decode(encoded []int, first int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(encoded)+1)
	result[0] = first
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(encoded); i++ {
		result[i+1] = result[i] ^ encoded[i]
	}
	return result
}

func main() {
	fmt.Println(Decode([]int{1, 2, 3}, 1))
	fmt.Println(Decode([]int{6, 2, 7, 3}, 4))
}
```
