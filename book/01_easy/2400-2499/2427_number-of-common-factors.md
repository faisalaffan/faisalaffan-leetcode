# 2427 — Number Of Common Factors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfCommonFactors(a int, b int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2427: Number of Common Factors
// https://leetcode.com/problems/number-of-common-factors/
// Difficulty: Easy
// Time O(min(a,b)) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfCommonFactors(12, 6)) // 4
	fmt.Println(NumberOfCommonFactors(25, 30)) // 2
}

func NumberOfCommonFactors(a int, b int) int {
	count := 0
	n := a
	if b < n {
		n = b
	}
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}
```
