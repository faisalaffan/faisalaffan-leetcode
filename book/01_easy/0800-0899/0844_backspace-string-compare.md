# 0844 — Backspace String Compare

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func backspaceCompare(s string, t string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #844: Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
	fmt.Println(backspaceCompare("a##c", "#a#c")) // true
}

// backspaceCompare checks if two strings are equal after applying backspace.
// Time: O(n + m). Space: O(1).
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 {
		// Find next valid char in s
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		// Find next valid char in t
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
```
