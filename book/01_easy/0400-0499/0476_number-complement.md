# 0476 — Number Complement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberComplement(num int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #476: Number Complement
// https://leetcode.com/problems/number-complement/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func NumberComplement(num int) int {
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	return ^num & ^mask
}

func main() {
	fmt.Println(NumberComplement(5))
	fmt.Println(NumberComplement(1))
	fmt.Println(NumberComplement(2))
}
```
