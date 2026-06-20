# 0387 — First Unique Character In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstUniqueCharacterInAString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #387: First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FirstUniqueCharacterInAString(s string) int {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	for i, c := range s {
		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstUniqueCharacterInAString("leetcode"))
	fmt.Println(FirstUniqueCharacterInAString("loveleetcode"))
	fmt.Println(FirstUniqueCharacterInAString("aabb"))
}
```
