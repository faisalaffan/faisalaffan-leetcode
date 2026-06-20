# 3438 — Find Valid Pair Of Adjacent Digits In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindValidPairOfAdjacentDigitsInString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[byte]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
  // Loop linear O(n): iterasi setiap elemen
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
