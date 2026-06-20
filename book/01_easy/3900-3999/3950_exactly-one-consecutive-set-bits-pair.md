# 3950 — Exactly One Consecutive Set Bits Pair

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ExactlyOneConsecutiveSetBitsPair(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3950: Exactly One Consecutive Set Bits Pair
// https://leetcode.com/problems/exactly-one-consecutive-set-bits-pair/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(6))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(5))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(3))
}

// Time: O(1)
// Space: O(1)
func ExactlyOneConsecutiveSetBitsPair(n int) bool {
	m := n & (n >> 1)
	return m > 0 && m&(m-1) == 0
}
```
