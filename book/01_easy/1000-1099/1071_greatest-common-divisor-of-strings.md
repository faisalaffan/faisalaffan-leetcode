# 1071 — Greatest Common Divisor Of Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func gcdOfStrings(str1, str2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Difficulty: Easy
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC")) // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
}

// LeetCode submission: gcdOfStrings
func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
