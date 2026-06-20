# 0737 — Sentence Similarity Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(n * alpha(n))  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #737: Sentence Similarity II
// https://leetcode.com/problems/sentence-similarity-ii/
// Difficulty: Medium [Paid]
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	pairs := [][]string{
		{"great", "fine"},
		{"drama", "acting"},
		{"fine", "good"},
	}
	fmt.Println(areSentencesSimilarTwo([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, pairs))
}

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	parent := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(words1); i++ {
		if words1[i] == words2[i] {
			continue
		}
		if find(words1[i]) != find(words2[i]) {
			return false
		}
	}

	return true
}
```
