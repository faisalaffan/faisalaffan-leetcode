# 1408 — String Matching In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func stringMatching(words []string) []string

import (
	"fmt"
	"sort"
	"strings"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 * L) where L is average word length, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1408: String Matching in an Array
// https://leetcode.com/problems/string-matching-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func stringMatching(words []string) []string

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(StringMatchingInAnArray([]string{"mass", "as", "hero", "superhero"})) // [as hero]
	fmt.Println(StringMatchingInAnArray([]string{"leetcode", "et", "code"}))          // [et code]
}

// Time: O(n^2 * L) where L is average word length, Space: O(n)
func StringMatchingInAnArray(words []string) []string {
  // Custom sort dengan comparator
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := make([]string, 0)
	for i, w := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], w) {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
```
