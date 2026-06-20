# 2820 — Election Results

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ElectionResults(votes []Vote) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2820: Election Results
// https://leetcode.com/problems/election-results/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Vote struct {
	Voter   string
	Candidate string
}

func ElectionResults(votes []Vote) string {
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[string]int)
	for _, v := range votes {
		counts[v.Candidate]++
	}

	type kv struct {
		Candidate string
		Count     int
	}
	sorted := make([]kv, 0, len(counts))
	for k, v := range counts {
		sorted = append(sorted, kv{k, v})
	}
  // Custom sort dengan comparator
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count != sorted[j].Count {
			return sorted[i].Count > sorted[j].Count
		}
		return sorted[i].Candidate < sorted[j].Candidate
	})

	if len(sorted) > 0 {
		return sorted[0].Candidate
	}
	return ""
}

func main() {
	votes := []Vote{
		{"A", "Alice"}, {"B", "Bob"}, {"C", "Alice"},
	}
	fmt.Println(ElectionResults(votes))

	votes2 := []Vote{
		{"X", "Cand1"}, {"Y", "Cand1"}, {"Z", "Cand2"},
	}
	fmt.Println(ElectionResults(votes2))
}
```
