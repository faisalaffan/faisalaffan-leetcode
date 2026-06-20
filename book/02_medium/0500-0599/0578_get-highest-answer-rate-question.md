# 0578 — Get Highest Answer Rate Question

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MostAnsweredQuestion(surveyLog [][]interface{}) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	shows := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
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
