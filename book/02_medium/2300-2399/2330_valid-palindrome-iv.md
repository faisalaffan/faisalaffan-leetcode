# 2330 — Valid Palindrome Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func makePalindrome(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2330: Valid Palindrome IV
// https://leetcode.com/problems/valid-palindrome-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func makePalindrome(s string) bool {
	diff := 0
	left, right := 0, len(s)-1

  // Two-pointer loop
	for left < right {
		if s[left] != s[right] {
			diff++
			if diff > 2 {
				return false
			}
		}
		left++
		right--
	}
	return diff <= 2
}

func main() {
	// Test case 1
	fmt.Println(makePalindrome("abcdba"))
	// Expected: true

	// Test case 2
	fmt.Println(makePalindrome("abccba"))
	// Expected: false (already palindrome, need exactly 2 changes)

	// Test case 3
	fmt.Println(makePalindrome("abcdeba"))
	// Expected: false (need 3 changes)
}
```
