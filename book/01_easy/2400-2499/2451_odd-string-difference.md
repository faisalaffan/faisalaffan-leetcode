# 2451 — Odd String Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func differenceArray(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2451: Odd String Difference
// https://leetcode.com/problems/odd-string-difference/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(OddStringDifference([]string{"adc", "wzy", "abc"})) // "abc"
	fmt.Println(OddStringDifference([]string{"aaa", "bob", "ccc", "ddd"})) // "bob"
}

func differenceArray(s string) string {
	diff := make([]byte, len(s)-1)
	for i := 1; i < len(s); i++ {
		diff[i-1] = s[i] - s[i-1]
	}
	return string(diff)
}

func OddStringDifference(words []string) string {
	diff0 := differenceArray(words[0])
	diff1 := differenceArray(words[1])

	if string(diff0) != string(diff1) {
		diff2 := differenceArray(words[2])
		if string(diff0) == string(diff2) {
			return words[1]
		}
		return words[0]
	}

	for i := 2; i < len(words); i++ {
		if string(differenceArray(words[i])) != string(diff0) {
			return words[i]
		}
	}
	return ""
}
```
