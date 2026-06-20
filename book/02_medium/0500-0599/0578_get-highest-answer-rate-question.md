# 0578 — Get Highest Answer Rate Question

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MostAnsweredQuestion(surveyLog [][]interface{}) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #578: Get Highest Answer Rate Question
// https://leetcode.com/problems/get-highest-answer-rate-question/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// survey_log: {question_id, action}
	// action: "show", "answer", "skip"
	surveyLog := [][]interface{}{
		{1, "show"},
		{1, "answer"},
		{2, "show"},
		{2, "skip"},
		{3, "show"},
		{3, "answer"},
	}
	fmt.Println(MostAnsweredQuestion(surveyLog))
}

func MostAnsweredQuestion(surveyLog [][]interface{}) int {
  // HashMap: O(1) lookup
	shows := make(map[int]int)
  // HashMap: O(1) lookup
	answers := make(map[int]int)

	for _, entry := range surveyLog {
		qID := entry[0].(int)
		action := entry[1].(string)
		shows[qID]++
		if action == "answer" {
			answers[qID]++
		}
	}

	bestQ := -1
	bestRate := -1.0

	for qID := range shows {
		rate := float64(answers[qID]) / float64(shows[qID])
		if rate > bestRate {
			bestRate = rate
			bestQ = qID
		}
	}

	return bestQ
}
```
