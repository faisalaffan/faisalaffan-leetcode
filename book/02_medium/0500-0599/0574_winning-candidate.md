# 0574 — Winning Candidate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindWinningCandidate(votes []int, candidateNames map[int]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
