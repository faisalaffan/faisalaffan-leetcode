# 3064 — Guess The Number Using Bitwise Questions I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func commonSetBits(num int, hidden int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Bitmask

**Waktu:** O(log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Bitmask** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3064: Guess the Number Using Bitwise Questions I (PAID)
// https://leetcode.com/problems/guess-the-number-using-bitwise-questions-i/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(1)

// Given a hidden number n (1 ≤ n < 2^30), determine n by calling the API
// commonSetBits(num int) int which returns popcount(n & num).
// For each bit position i, probe commonSetBits(1<<i) — if the result is > 0,
// that bit is set in n.

import "fmt"

// commonSetBits simulates the LeetCode API: popcount of (num & hidden).
func commonSetBits(num int, hidden int) int {
	and := hidden & num
	cnt := 0
	for and > 0 {
		cnt++
		and &= and - 1
	}
	return cnt
}

func main() {
	fmt.Println(findNumber(31)) // 31 (0b11111)
	fmt.Println(findNumber(33)) // 33 (0b100001)
	fmt.Println(findNumber(1))  // 1
	fmt.Println(findNumber(0))  // 0 (edge case, though problem says n >= 1)
}

func findNumber(n int) int {
	result := 0
	for i := 0; i < 31; i++ {
		if commonSetBits(1<<i, n) > 0 {
			result |= 1 << i
		}
	}
	return result
}
```
