# 0602 — Friend Requests Ii Who Has The Most Friends

## Deskripsi

**Soal:** [0602. Friend Requests Ii Who Has The Most Friends](https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #602: Friend Requests II: Who Has the Most Friends
// https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Friend requests: {requester_id, accepter_id}
	requests := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	fmt.Println(MostFriends(requests))
}

func MostFriends(requests [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	friendCount := make(map[int]int)
	for _, req := range requests {
		friendCount[req[0]]++
		friendCount[req[1]]++
	}

	maxCount := 0
	maxID := 0
	for id, count := range friendCount {
		if count > maxCount {
			maxCount = count
			maxID = id
		}
	}

	return maxID
}
```
