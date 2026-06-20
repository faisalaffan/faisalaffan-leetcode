# 3340 — Check Balanced String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckBalancedString(num string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3340: Check Balanced String
// https://leetcode.com/problems/check-balanced-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckBalancedString("1234"))
	fmt.Println(CheckBalancedString("24123"))
}

// CheckBalancedString returns true if sum of digits at even positions equals sum at odd positions.
// Time: O(n). Space: O(1).
func CheckBalancedString(num string) bool {
	evenSum, oddSum := 0, 0
	for i, ch := range num {
		digit := int(ch - '0')
		if i%2 == 0 {
			evenSum += digit
		} else {
			oddSum += digit
		}
	}
	return evenSum == oddSum
}
```
