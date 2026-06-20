# 2452 — Words Within Two Edits Of Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func twoEditWords(queries []string, dictionary []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m * L)  
**Kompleksitas Ruang:** O(1) where L = word length

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2452: Words Within Two Edits of Dictionary
// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// Difficulty: Medium
// Time: O(n * m * L) | Space: O(1) where L = word length
// For each query word, check Hamming distance to each dictionary word.

import "fmt"

func main() {
	fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})) // ["word","note","wood"]
	fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))                                            // []
}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	for _, q := range queries {
		for _, d := range dictionary {
			if hammingDist(q, d) <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

func hammingDist(a, b string) int {
	dist := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
```
