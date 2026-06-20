# 0869 — Reordered Power Of 2

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ReorderedPowerOfTwo(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #869: Reordered Power of 2
// https://leetcode.com/problems/reordered-power-of-2/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ReorderedPowerOfTwo(1))
	fmt.Println(ReorderedPowerOfTwo(10))
	fmt.Println(ReorderedPowerOfTwo(46))
}

// Time: O(log n) | Space: O(1)
func ReorderedPowerOfTwo(n int) bool {
	sig := signature(n)
	for i := 1; i <= 1_000_000_000; i <<= 1 {
		if signature(i) == sig {
			return true
		}
	}
	return false
}

func signature(x int) [10]int {
	var cnt [10]int
	for x > 0 {
		cnt[x%10]++
		x /= 10
	}
	return cnt
}
```
