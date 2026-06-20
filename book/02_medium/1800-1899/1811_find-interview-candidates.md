# 1811 — Find Interview Candidates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findCandidates(contests []ContestScore, submissions []UserContest) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1811: Find Interview Candidates
// https://leetcode.com/problems/find-interview-candidates/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type ContestScore struct {
	ContestID int
	GoldID    int // user ID of gold medalist
	SilverID  int
	BronzeID  int
}

type UserContest struct {
	UserID    int
	ContestID int
	Score     int
}

func findCandidates(contests []ContestScore, submissions []UserContest) []int {
  // Membuat map (HashMap) — pencarian O(1)
	medalCount := make(map[int]int)
	for _, c := range contests {
		medalCount[c.GoldID]++
		medalCount[c.SilverID]++
		medalCount[c.BronzeID]++
	}

  // Membuat map (HashMap) — pencarian O(1)
	candidateSet := make(map[int]bool)

	// Any user with 3+ medals
	for user, count := range medalCount {
		if count >= 3 {
			candidateSet[user] = true
		}
	}

	// Users who won gold in consecutive contests (contests can be consecutive by contest_id)
  // Membuat map (HashMap) — pencarian O(1)
	userContests := make(map[int][]int)
	for _, c := range contests {
		userContests[c.GoldID] = append(userContests[c.GoldID], c.ContestID)
	}

	for user, ids := range userContests {
		// Sort contest IDs
		for i := 1; i < len(ids); i++ {
			if ids[i] == ids[i-1]+1 {
				candidateSet[user] = true
				break
			}
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, len(candidateSet))
	for u := range candidateSet {
		result = append(result, u)
	}
	return result
}

func main() {
	contests := []ContestScore{
		{1, 1, 2, 3},
		{2, 1, 4, 5},
		{3, 1, 6, 7},
	}
	fmt.Println(findCandidates(contests, nil)) // Expected: [1] (gold in 3 consecutive contests)
}
```
