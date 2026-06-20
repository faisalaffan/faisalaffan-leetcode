# 0242 — Valid Anagram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func IsAnagram(s string, t string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) (fixed 26 chars)


## 💻 Solusi Go

```go
package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram("rat", "car"))
}
```
