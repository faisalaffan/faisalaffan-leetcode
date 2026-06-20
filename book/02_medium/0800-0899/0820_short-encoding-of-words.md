# 0820 — Short Encoding Of Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortEncodingOfWords(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L^2)  
**Kompleksitas Ruang:** O(n * L) where L is average word length

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #820: Short Encoding of Words
// https://leetcode.com/problems/short-encoding-of-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShortEncodingOfWords([]string{"time", "me", "bell"}))
	fmt.Println(ShortEncodingOfWords([]string{"t"}))
	fmt.Println(ShortEncodingOfWords([]string{"me", "time"}))
}

// Time: O(n * L^2) | Space: O(n * L) where L is average word length
func ShortEncodingOfWords(words []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(set, word[i:])
		}
	}

	ans := 0
	for word := range set {
		ans += len(word) + 1
	}
	return ans
}
```
