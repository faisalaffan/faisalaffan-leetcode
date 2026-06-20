# 0634 — Find The Derangement Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindDerangement(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #634: Find the Derangement of An Array
// https://leetcode.com/problems/find-the-derangement-of-an-array/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindDerangement(3))
	fmt.Println(FindDerangement(4))
}

func FindDerangement(n int) int {
  // Edge case: input kosong
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	const mod = 1_000_000_007
	a, b := 0, 1 // D(1)=0, D(2)=1

	for i := 3; i <= n; i++ {
		c := ((i - 1) * (a + b)) % mod
		a, b = b, c
	}

	return b
}
```
