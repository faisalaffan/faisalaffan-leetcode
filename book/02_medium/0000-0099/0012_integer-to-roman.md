# 0012 — Integer To Roman

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi angka Romawi. Tugasmu adalah mengkonversinya ke integer.

Simbol: I=1, V=5, X=10, L=50, C=100, D=500, M=1000. Aturan: simbol kecil di DEPAN simbol besar → kurangi (IV=4). Simbol kecil di BELAKANG → tambah (VI=6).

**Cara berpikir:** Iterasi dari kiri ke kanan. Kalau nilai saat ini < nilai berikutnya → kurangi. Selain itu → tambahkan.

**Fungsi Solusi:** `func intToRoman(num int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #12: Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
// Difficulty: Medium

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
  // Linear scan O(n)
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(intToRoman(3749)) // "MMMDCCXLIX"

	// Test case 2
	fmt.Println(intToRoman(58)) // "LVIII"

	// Test case 3
	fmt.Println(intToRoman(1994)) // "MCMXCIV"
}

// Time: O(1) | Space: O(1)
```
