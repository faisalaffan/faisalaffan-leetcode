# 1355 — Activity Participants

## Deskripsi

**Soal:** [1355. Activity Participants](https://leetcode.com/problems/activity-participants/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of friends  
**Kompleksitas Ruang:** O(m) where m = number of activities

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
