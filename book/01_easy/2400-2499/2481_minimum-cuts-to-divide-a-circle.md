# 2481 — Minimum Cuts To Divide A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumCutsToDivideACircle(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2481: Minimum Cuts to Divide a Circle
// https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCutsToDivideACircle(4)) // 2
	fmt.Println(MinimumCutsToDivideACircle(3)) // 3
	fmt.Println(MinimumCutsToDivideACircle(1)) // 0
}

func MinimumCutsToDivideACircle(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}
```
