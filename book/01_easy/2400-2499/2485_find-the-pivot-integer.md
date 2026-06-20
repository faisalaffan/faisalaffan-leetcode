# 2485 — Find The Pivot Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindThePivotInteger(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2485: Find the Pivot Integer
// https://leetcode.com/problems/find-the-pivot-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindThePivotInteger(8)) // 6
	fmt.Println(FindThePivotInteger(1)) // 1
	fmt.Println(FindThePivotInteger(4)) // -1
}

func FindThePivotInteger(n int) int {
	total := n * (n + 1) / 2
	sum := 0
	for x := 1; x <= n; x++ {
		sum += x
		if sum == total-sum+x {
			return x
		}
	}
	return -1
}
```
