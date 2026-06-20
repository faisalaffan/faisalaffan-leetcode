# 2330 — Valid Palindrome Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func makePalindrome(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2330: Valid Palindrome IV
// https://leetcode.com/problems/valid-palindrome-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func makePalindrome(s string) bool {
	diff := 0
	left, right := 0, len(s)-1

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		if s[left] != s[right] {
			diff++
			if diff > 2 {
				return false
			}
		}
		left++
		right--
	}
	return diff <= 2
}

func main() {
	// Test case 1
	fmt.Println(makePalindrome("abcdba"))
	// Expected: true

	// Test case 2
	fmt.Println(makePalindrome("abccba"))
	// Expected: false (already palindrome, need exactly 2 changes)

	// Test case 3
	fmt.Println(makePalindrome("abcdeba"))
	// Expected: false (need 3 changes)
}
```
