# 0678 — Valid Parenthesis String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkValidString(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #678: Valid Parenthesis String
// https://leetcode.com/problems/valid-parenthesis-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
	low, high := 0, 0

	for _, c := range s {
		if c == '(' {
			low++
			high++
		} else if c == ')' {
			low--
			high--
		} else { // '*'
			low--
			high++
		}

		if high < 0 {
			return false
		}
		if low < 0 {
			low = 0
		}
	}

	return low == 0
}
```
