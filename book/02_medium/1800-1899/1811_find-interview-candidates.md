# 1811 — Find Interview Candidates

## Deskripsi

**Soal:** [1811. Find Interview Candidates](https://leetcode.com/problems/find-interview-candidates/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func findCandidates(contests []ContestScore, submissions []UserContest) []int`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	medalCount := make(map[int]int)
	for _, c := range contests {
		medalCount[c.GoldID]++
		medalCount[c.SilverID]++
		medalCount[c.BronzeID]++
	}

  // Membuat map untuk pencarian O(1): key → value
	candidateSet := make(map[int]bool)

	// Any user with 3+ medals
	for user, count := range medalCount {
		if count >= 3 {
			candidateSet[user] = true
		}
	}

	// Users who won gold in consecutive contests (contests can be consecutive by contest_id)
  // Membuat map untuk pencarian O(1): key → value
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

  // Membuat slice untuk menyimpan hasil
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
