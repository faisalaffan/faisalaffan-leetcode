# 0784 — Letter Case Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func letterCasePermutation(s string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(2^n * n)  
**Kompleksitas Ruang:** O(2^n * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #784: Letter Case Permutation
// https://leetcode.com/problems/letter-case-permutation/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	curr := make([]byte, len(s))

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			result = append(result, string(curr))
			return
		}

		curr[idx] = s[idx]
		if unicode.IsLetter(rune(s[idx])) {
			curr[idx] = byte(unicode.ToLower(rune(s[idx])))
			backtrack(idx + 1)
			curr[idx] = byte(unicode.ToUpper(rune(s[idx])))
			backtrack(idx + 1)
		} else {
			backtrack(idx + 1)
		}
	}

	backtrack(0)
	return result
}
```
