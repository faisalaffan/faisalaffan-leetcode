# 0089 — Gray Code

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func grayCode(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(2^n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #89: Gray Code
// https://leetcode.com/problems/gray-code/
// Difficulty: Medium

import "fmt"

func grayCode(n int) []int {
  // Alokasi slice
	result := make([]int, 1<<n)
  // Linear scan O(n)
	for i := 0; i < len(result); i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(grayCode(2)) // [0 1 3 2]

	// Test case 2
	fmt.Println(grayCode(1)) // [0 1]

	// Test case 3
	fmt.Println(grayCode(3)) // [0 1 3 2 6 7 5 4]
}

// Time: O(2^n) | Space: O(1)
```
