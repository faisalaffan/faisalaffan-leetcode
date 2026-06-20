# 0178 — Rank Scores

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rankScores(scores []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #178: Rank Scores
// https://leetcode.com/problems/rank-scores/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func rankScores(scores []int) []int {
	if len(scores) == 0 {
		return nil
	}

	type pair struct {
		score int
		idx   int
	}

	pairs := make([]pair, len(scores))
	for i, s := range scores {
		pairs[i] = pair{s, i}
	}

  // Custom sort dengan comparator
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

  // Alokasi slice integer
	ranks := make([]int, len(scores))
	rank := 1
	for i, p := range pairs {
		if i > 0 && p.score < pairs[i-1].score {
			rank = i + 1
		}
		ranks[p.idx] = rank
	}

	return ranks
}

func main() {
	fmt.Println(rankScores([]int{100, 90, 90, 80}))
	fmt.Println(rankScores([]int{50, 60, 70}))
	fmt.Println(rankScores([]int{90, 80, 80, 70, 60, 60}))
}
```
