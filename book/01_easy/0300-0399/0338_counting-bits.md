# 0338 — Counting Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountingBits(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #338: Counting Bits
// https://leetcode.com/problems/counting-bits/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CountingBits(n int) []int {
  // Alokasi slice
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	fmt.Println(CountingBits(2))
	fmt.Println(CountingBits(5))
	fmt.Println(CountingBits(0))
}
```
