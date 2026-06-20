# 3058 — Friends With No Mutual Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func friendsWithNoMutualFriends(friendships []Friendship) []NoMutualPair
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * d^2) where d is avg degree  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3058: Friends With No Mutual Friends
// https://leetcode.com/problems/friends-with-no-mutual-friends/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n * d^2) where d is avg degree | Space: O(n)

import (
	"fmt"
	"sort"
)

type Friendship struct {
	UserID1 int
	UserID2 int
}

type NoMutualPair struct {
	UserID1 int
	UserID2 int
}

func friendsWithNoMutualFriends(friendships []Friendship) []NoMutualPair {
	// Build adjacency set: user -> set of friends
  // Membuat map (HashMap) — pencarian O(1)
	adj := make(map[int]map[int]bool)
	for _, f := range friendships {
		if adj[f.UserID1] == nil {
			adj[f.UserID1] = make(map[int]bool)
		}
		if adj[f.UserID2] == nil {
			adj[f.UserID2] = make(map[int]bool)
		}
		adj[f.UserID1][f.UserID2] = true
		adj[f.UserID2][f.UserID1] = true
	}

	var results []NoMutualPair
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[[2]int]bool)

	for _, f := range friendships {
		a, b := f.UserID1, f.UserID2

		// Ensure we only process each unordered pair once
		pair := [2]int{a, b}
		if seen[pair] {
			continue
		}
		seen[pair] = true
		seen[[2]int{b, a}] = true

		// Check if a and b have any mutual friends
		hasMutual := false
		for friend := range adj[a] {
			if friend == b {
				continue
			}
			if adj[b][friend] {
				hasMutual = true
				break
			}
		}

		if !hasMutual {
			results = append(results, NoMutualPair{UserID1: a, UserID2: b})
		}
	}

	// Order by user_id1 ASC, user_id2 ASC
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		if results[i].UserID1 != results[j].UserID1 {
			return results[i].UserID1 < results[j].UserID1
		}
		return results[i].UserID2 < results[j].UserID2
	})

	return results
}

func main() {
	// Test from problem description
	friendships := []Friendship{
		{UserID1: 1, UserID2: 2},
		{UserID1: 2, UserID2: 3},
		{UserID1: 2, UserID2: 4},
		{UserID1: 1, UserID2: 3},
		{UserID1: 3, UserID2: 4},
		{UserID1: 6, UserID2: 7},
		{UserID1: 8, UserID2: 9},
		{UserID1: 5, UserID2: 6},
	}

	fmt.Println("Friends With No Mutual Friends")
	fmt.Println("=============================")
	fmt.Printf("%-8s %s\n", "user1", "user2")
	fmt.Println("-------------------")

	results := friendsWithNoMutualFriends(friendships)
	for _, r := range results {
		fmt.Printf("%-8d %d\n", r.UserID1, r.UserID2)
	}
}
```
