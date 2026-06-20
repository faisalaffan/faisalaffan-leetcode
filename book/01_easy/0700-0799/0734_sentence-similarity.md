# 0734 — Sentence Similarity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + p) where n = len(sentence), p = len(pairs). Space: O(p).  
**Kompleksitas Ruang:** O(p).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #734: Sentence Similarity
// https://leetcode.com/problems/sentence-similarity/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	s1 := []string{"great", "acting", "skills"}
	s2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // true

	s1 = []string{"great"}
	s2 = []string{"great"}
	fmt.Println(areSentencesSimilar(s1, s2, [][]string{})) // true

	s1 = []string{"great"}
	s2 = []string{"doubleplus", "good"}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // false
}

// areSentencesSimilar checks if two sentences are similar according to given similarity pairs.
// Time: O(n + p) where n = len(sentence), p = len(pairs). Space: O(p).
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	// Build bidirectional map
  // Membuat map (HashMap) — pencarian O(1)
	pairMap := make(map[string]map[string]bool)
	for _, p := range similarPairs {
		a, b := p[0], p[1]
		if pairMap[a] == nil {
			pairMap[a] = make(map[string]bool)
		}
		if pairMap[b] == nil {
			pairMap[b] = make(map[string]bool)
		}
		pairMap[a][b] = true
		pairMap[b][a] = true
	}

  // Range loop: iterasi dengan indeks + nilai
	for i := range sentence1 {
		w1, w2 := sentence1[i], sentence2[i]
		if w1 == w2 {
			continue
		}
		if !pairMap[w1][w2] {
			return false
		}
	}
	return true
}
```
