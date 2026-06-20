# 2562 — Find The Array Concatenation Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheArrayConcatenationValue(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2562: Find the Array Concatenation Value
// https://leetcode.com/problems/find-the-array-concatenation-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheArrayConcatenationValue([]int{7, 52, 2, 4})) // 596
	fmt.Println(FindTheArrayConcatenationValue([]int{5, 14, 13, 8, 12})) // 673
}

func FindTheArrayConcatenationValue(nums []int) int64 {
	sum := int64(0)
	i, j := 0, len(nums)-1
	for i < j {
		concat := int64(nums[i])
		n := nums[j]
		digits := 0
  // Edge case: input kosong
		if n == 0 {
			digits = 1
		} else {
			digits = int(math.Log10(float64(n))) + 1
		}
		concat = concat * pow10(digits) + int64(n)
		sum += concat
		i++
		j--
	}
	if i == j {
		sum += int64(nums[i])
	}
	return sum
}

func pow10(n int) int64 {
	p := int64(1)
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}
```
