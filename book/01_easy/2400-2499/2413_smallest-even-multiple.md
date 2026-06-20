# 2413 — Smallest Even Multiple

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SmallestEvenMultiple(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2413: Smallest Even Multiple
// https://leetcode.com/problems/smallest-even-multiple/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(SmallestEvenMultiple(5)) // 10
	fmt.Println(SmallestEvenMultiple(6)) // 6
}

func SmallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
```
