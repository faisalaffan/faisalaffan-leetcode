# 2299 — Strong Password Checker Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func StrongPasswordCheckerIi(password string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2299: Strong Password Checker II
// https://leetcode.com/problems/strong-password-checker-ii/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(StrongPasswordCheckerIi("IloveLe3tcode!")) // true
	fmt.Println(StrongPasswordCheckerIi("Me+You--IsMyDream")) // false
	fmt.Println(StrongPasswordCheckerIi("1aB!")) // false
}

func StrongPasswordCheckerIi(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false
	special := "!@#$%^&*()-+"

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(password); i++ {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		ch := password[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= '0' && ch <= '9':
			hasDigit = true
		default:
			for j := 0; j < len(special); j++ {
				if ch == special[j] {
					hasSpecial = true
					break
				}
			}
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}
```
