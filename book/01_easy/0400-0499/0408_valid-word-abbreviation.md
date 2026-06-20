# 0408 — Valid Word Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidWordAbbreviation(word, abbr string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #408: Valid Word Abbreviation
// https://leetcode.com/problems/valid-word-abbreviation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(1)
func ValidWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
			continue
		}
		if abbr[j] == '0' {
			return false
		}
		num := 0
		for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
			num = num*10 + int(abbr[j]-'0')
			j++
		}
		i += num
	}
	return i == len(word) && j == len(abbr)
}

func main() {
	fmt.Println(ValidWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(ValidWordAbbreviation("apple", "a2e"))
	fmt.Println(ValidWordAbbreviation("hi", "1"))
}
```
