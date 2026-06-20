# 2000 — Reverse Prefix Of Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReversePrefixOfWord(word string, ch byte) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2000: Reverse Prefix of Word
// https://leetcode.com/problems/reverse-prefix-of-word/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReversePrefixOfWord("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(ReversePrefixOfWord("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(ReversePrefixOfWord("abcd", 'z'))    // "abcd"
}

// Time: O(n), Space: O(n)
func ReversePrefixOfWord(word string, ch byte) string {
	idx := -1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}

	result := make([]byte, len(word))
	for i := 0; i <= idx; i++ {
		result[i] = word[idx-i]
	}
	for i := idx + 1; i < len(word); i++ {
		result[i] = word[i]
	}
	return string(result)
}
```
