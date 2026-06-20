# 1945 — Sum Of Digits Of String After Convert

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfDigitsOfStringAfterConvert(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1945: Sum of Digits of String After Convert
// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfDigitsOfStringAfterConvert("iiii", 1))  // 36
	fmt.Println(SumOfDigitsOfStringAfterConvert("leetcode", 2))  // 6
}

// Time: O(n), Space: O(n)
func SumOfDigitsOfStringAfterConvert(s string, k int) int {
	var digits string
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		digits += strconv.Itoa(int(s[i] - 'a' + 1))
	}

	for t := 0; t < k; t++ {
		sum := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(digits); i++ {
			sum += int(digits[i] - '0')
		}
		digits = strconv.Itoa(sum)
	}

	result, _ := strconv.Atoi(digits)
	return result
}
```
