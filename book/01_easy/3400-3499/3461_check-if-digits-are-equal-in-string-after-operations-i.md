# 3461 — Check If Digits Are Equal In String After Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3461: Check If Digits Are Equal in String After Operations I
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1234"))
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1111"))
}

// CheckIfDigitsAreEqualInStringAfterOperationsI repeatedly replaces adjacent digit pairs with (sum % 10) until 2 digits remain, then checks equality.
// Time: O(n^2). Space: O(n).
func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool {
  // Alokasi slice integer
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	for len(digits) > 2 {
  // Alokasi slice integer
		next := make([]int, len(digits)-1)
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(digits)-1; i++ {
			next[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = next
	}
	return digits[0] == digits[1]
}
```
