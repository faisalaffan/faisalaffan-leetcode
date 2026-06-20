# 0422 — Valid Word Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidWordSquare(words []string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #422: Valid Word Square
// https://leetcode.com/problems/valid-word-square/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n*m), Space: O(1)
func ValidWordSquare(words []string) bool {
  // Linear scan O(n)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crmy", "dtyx"}))
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(ValidWordSquare([]string{"ball", "area", "lead", "lady"}))
}
```
