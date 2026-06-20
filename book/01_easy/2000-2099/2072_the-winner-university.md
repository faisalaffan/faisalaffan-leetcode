# 2072 — The Winner University

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func TheWinnerUniversity(newYork []int, california []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2072: The Winner University
// https://leetcode.com/problems/the-winner-university/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// New York students: (score)
	newYork := []int{90, 85, 78}
	// California students: (score)
	california := []int{92, 88, 70}
	fmt.Println(TheWinnerUniversity(newYork, california)) // "California University"
}

// Time: O(n), Space: O(1)
func TheWinnerUniversity(newYork []int, california []int) string {
	nyScore := 0
	for _, s := range newYork {
		nyScore += s
	}
	caScore := 0
	for _, s := range california {
		caScore += s
	}

	if nyScore > caScore {
		return "New York University"
	} else if caScore > nyScore {
		return "California University"
	}
	return "No Winner"
}
```
