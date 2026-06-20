# 1355 — Activity Participants

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func activityParticipants(activities []struct { id int name string }, friends []struct { id int name string activityID int }) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = number of friends  |  **Ruang:** O(m) where m = number of activities

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1355: Activity Participants
// https://leetcode.com/problems/activity-participants/
// Difficulty: Medium

import "fmt"

func main() {
	activities := []struct {
		id   int
		name string
	}{
		{1, "Eating"},
		{2, "Singing"},
		{3, "Horse Riding"},
	}
	friends := []struct {
		id         int
		name       string
		activityID int
	}{
		{1, "Jonathan D.", 1},
		{2, "Jade W.", 1},
		{3, "Victor J.", 1},
		{4, "Elvis O.", 2},
		{5, "Daniel A.", 2},
		{6, "Bob B.", 3},
	}

	result := activityParticipants(activities, friends)
	fmt.Println(result) // ["Singing"]
}

// Time: O(n) where n = number of friends
// Space: O(m) where m = number of activities
func activityParticipants(activities []struct {
	id   int
	name string
}, friends []struct {
	id         int
	name       string
	activityID int
}) []string {
  // HashMap: O(1) lookup
	counts := make(map[int]int)
	for _, f := range friends {
		counts[f.activityID]++
	}

	if len(counts) == 0 {
		return nil
	}

	// Find min and max counts
	minCount, maxCount := len(friends), 0
	for _, c := range counts {
		if c < minCount {
			minCount = c
		}
		if c > maxCount {
			maxCount = c
		}
	}

	// Find activities with count between min and max (non-inclusive)
	var result []string
	for _, a := range activities {
		c := counts[a.id]
		if c > minCount && c < maxCount {
			result = append(result, a.name)
		}
	}
	return result
}
```
