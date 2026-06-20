# 1957 — Delete Characters To Make Fancy String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func DeleteCharactersToMakeFancyString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1957: Delete Characters to Make Fancy String
// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DeleteCharactersToMakeFancyString("leeetcode"))     // "leetcode"
	fmt.Println(DeleteCharactersToMakeFancyString("aaabaaaa"))      // "aabaa"
	fmt.Println(DeleteCharactersToMakeFancyString("aab"))           // "aab"
}

// Time: O(n), Space: O(n)
func DeleteCharactersToMakeFancyString(s string) string {
	result := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n >= 2 && result[n-1] == s[i] && result[n-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
```
