# 2802 — Find The K Th Lucky Number

## Deskripsi

**Soal:** [2802. Find The K Th Lucky Number](https://leetcode.com/problems/find-the-k-th-lucky-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log k)  
**Kompleksitas Ruang:** O(log k)

**Algoritma:** —

**Fungsi Solusi:** `func FindTheKThLuckyNumber(k int) string`

## Solusi Go

```go
package main

// LeetCode #2802: Find The K-th Lucky Number
// https://leetcode.com/problems/find-the-k-th-lucky-number/
// Difficulty: Medium [Paid]
// Time: O(log k) | Space: O(log k)

import "fmt"

func FindTheKThLuckyNumber(k int) string {
	// K-th lucky number: binary representation of k+1, then replace 0->4, 1->7
	// n = k + 1
	n := k + 1
	binary := fmt.Sprintf("%b", n)
	// Remove first '1' (it's the leading bit from the offset)
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, len(binary)-1)
	for i := 1; i < len(binary); i++ {
		if binary[i] == '0' {
			result[i-1] = '4'
		} else {
			result[i-1] = '7'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FindTheKThLuckyNumber(1))
	fmt.Println(FindTheKThLuckyNumber(3))
	fmt.Println(FindTheKThLuckyNumber(5))
}
```
