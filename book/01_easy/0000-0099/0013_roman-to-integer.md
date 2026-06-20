# 0013 — Roman To Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi angka Romawi. Tugasmu adalah mengkonversinya ke integer.

Simbol: I=1, V=5, X=10, L=50, C=100, D=500, M=1000. Aturan: simbol kecil di DEPAN simbol besar → kurangi (IV=4). Simbol kecil di BELAKANG → tambah (VI=6).

**Cara berpikir:** Iterasi dari kiri ke kanan. Kalau nilai saat ini < nilai berikutnya → kurangi. Selain itu → tambahkan.

**Fungsi Solusi:** `func RomanToInt(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #13: Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RomanToInt(s string) int {
	vals := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	sum, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := vals[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur
	}
	return sum
}

func main() {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))
}
```
