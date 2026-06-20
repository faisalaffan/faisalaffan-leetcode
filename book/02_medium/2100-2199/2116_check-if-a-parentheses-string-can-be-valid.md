# 2116 — Check If A Parentheses String Can Be Valid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func canBeValid(s string, locked string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2116: Check if a Parentheses String Can Be Valid
// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	// Left to right: check for excess ')'
	balance := 0
	flexible := 0
	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	// Right to left: check for excess '('
	balance = 0
	flexible = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == ')' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canBeValid("))()))", "010100"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canBeValid("()()", "0000"))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", canBeValid(")", "0"))
	// Expected: false (odd length)
}
```
