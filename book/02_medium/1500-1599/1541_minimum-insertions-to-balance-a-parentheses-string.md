# 1541 — Minimum Insertions To Balance A Parentheses String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinInsertions(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1541: Minimum Insertions to Balance a Parentheses String
// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinInsertions("(()))"))
	fmt.Println(MinInsertions("())"))
	fmt.Println(MinInsertions("))())("))
}

func MinInsertions(s string) int {
	// Time: O(N), Space: O(1)
	// Each '(' needs two ')' to balance.
	insertions := 0
	open := 0 // number of '(' that need closing

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else { // ')'
			if open > 0 {
				// Check if next char is also ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Both ')' found, consume both
					i++ // skip next ')'
				} else {
					// Need one more ')'
					insertions++
				}
				open--
			} else {
				// Need a '(' before this ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Insert '(' and consume both ')'
					insertions++ // for '('
					i++          // consume next ')'
				} else {
					// Insert '(' and one ')'
					insertions += 2
				}
			}
		}
	}

	// Each remaining '(' needs two ')'
	insertions += open * 2

	return insertions
}
```
