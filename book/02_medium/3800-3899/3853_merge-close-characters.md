# 3853 — Merge Close Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeCloseCharacters(s string, k int) string
```

> **💡 Hint:** Simulate merging. Track last position of each character.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N^2)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3853: Merge Close Characters
// https://leetcode.com/problems/merge-close-characters/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Simulate merging. Track last position of each character.
// When a char appears within distance k of its last occurrence, skip it.

import "fmt"

func MergeCloseCharacters(s string, k int) string {
	result := make([]byte, 0)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// Check if same char exists in result within distance k
		merge := false
		for j := len(result) - 1; j >= 0 && len(result)-1-j < k; j-- {
			if result[j] == ch {
				merge = true
				break
			}
		}
		if !merge {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	// Example 1
	fmt.Println(MergeCloseCharacters("abca", 3)) // Expected: "abc"

	// Example 2
	fmt.Println(MergeCloseCharacters("aabca", 2)) // Expected: "abca"

	// Example 3
	fmt.Println(MergeCloseCharacters("yybyzybz", 2)) // Expected: "ybzybz"
}
```
