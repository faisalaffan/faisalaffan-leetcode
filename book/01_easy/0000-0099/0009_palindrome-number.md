# 0009 — Palindrome Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah bilangan bulat (integer). Tugasmu adalah menentukan apakah bilangan tersebut adalah **palindrome** — dibaca sama dari depan maupun dari belakang.

Contoh: `121` → palindrome. `-121` → bukan (tanda minus!). `10` → bukan.

**Cara berpikir:** Balik setengah digit secara matematika menggunakan modulo (`%`) dan pembagian (`/`). Tidak perlu konversi ke string.

**Fungsi Solusi:** `func IsPalindrome(x int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #9: Palindrome Number
// https://leetcode.com/problems/palindrome-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return x == reverted || x == reverted/10
}

func main() {
	fmt.Println(IsPalindrome(121))  // true
	fmt.Println(IsPalindrome(-121)) // false
	fmt.Println(IsPalindrome(10))   // false
}
```
