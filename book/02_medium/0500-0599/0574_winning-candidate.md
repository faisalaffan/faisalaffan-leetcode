# 0574 — Winning Candidate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindWinningCandidate(votes []int, candidateNames map[int]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #574: Winning Candidate
// https://leetcode.com/problems/winning-candidate/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// votes: {candidate_id}
	votes := []int{1, 2, 2, 3, 3, 3}
	// candidate names: map candidate_id -> name
	candidateNames := map[int]int{1: 1001, 2: 1002, 3: 1003}
	fmt.Println(FindWinningCandidate(votes, candidateNames))
}

func FindWinningCandidate(votes []int, candidateNames map[int]int) int {
  // HashMap: O(1) lookup
	counts := make(map[int]int)
	for _, v := range votes {
		counts[v]++
	}

	maxVotes := 0
	winner := -1
	for id, count := range counts {
		if count > maxVotes {
			maxVotes = count
			winner = candidateNames[id]
		}
	}

	return winner
}
```
