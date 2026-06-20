# 1374 — Generate A String With Characters That Have Odd Counts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func generateTheString(n int) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1374: Generate a String With Characters That Have Odd Counts
// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
// Difficulty: Easy
//
// LeetCode submission: func generateTheString(n int) string

import "fmt"

func main() {
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(4)) // "aaab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(2)) // "ab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(7)) // "aaaaaaa"
}

// Time: O(n), Space: O(n)
func GenerateAStringWithCharactersThatHaveOddCounts(n int) string {
	if n%2 == 1 {
		return string(makeN('a', n))
	}
	return string(makeN('a', n-1)) + "b"
}

func makeN(ch byte, n int) []byte {
	res := make([]byte, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range res {
		res[i] = ch
	}
	return res
}
```
