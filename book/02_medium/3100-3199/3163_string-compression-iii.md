# 3163 — String Compression Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func compressedString(word string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3163: String Compression III
// https://leetcode.com/problems/string-compression-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func compressedString(word string) string {
	ans := make([]byte, 0, len(word)*2)
	n := len(word)
	i := 0

	for i < n {
		ch := word[i]
		j := i
		for j < n && j-i < 9 && word[j] == ch {
			j++
		}
		ans = append(ans, byte('0'+j-i), ch)
		i = j
	}
	return string(ans)
}

func main() {
	fmt.Println(compressedString("abcde"))               // Expected: "1a1b1c1d1e"
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))     // Expected: "9a5a2b"
}
```
