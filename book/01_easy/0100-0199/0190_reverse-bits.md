# 0190 — Reverse Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ReverseBits(num uint32) uint32`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #190: Reverse Bits
// https://leetcode.com/problems/reverse-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}
	return result
}

func main() {
	fmt.Println(ReverseBits(43261596))    // 964176192
	fmt.Println(ReverseBits(4294967293))  // 3221225471
}
```
