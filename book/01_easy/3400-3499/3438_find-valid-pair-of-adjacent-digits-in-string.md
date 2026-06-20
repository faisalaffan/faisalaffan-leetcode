# 3438 — Find Valid Pair Of Adjacent Digits In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func FindValidPairOfAdjacentDigitsInString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3438: Find Valid Pair of Adjacent Digits in String
// https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindValidPairOfAdjacentDigitsInString("2523533"))
	fmt.Println(FindValidPairOfAdjacentDigitsInString("111"))
}

// FindValidPairOfAdjacentDigitsInString finds the first pair of adjacent equal digits where the digit's frequency > the digit.
// Time: O(n). Space: O(n).
func FindValidPairOfAdjacentDigitsInString(s string) string {
  // HashMap: O(1) lookup
	freq := make(map[byte]int)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
  // Linear scan O(n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt := freq[s[i]]
			if cnt > int(s[i]-'0') {
				return s[i : i+2]
			}
		}
	}
	return ""
}
```
