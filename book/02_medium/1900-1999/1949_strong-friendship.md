# 1949 — Strong Friendship

## Deskripsi

**Soal:** [1949. Strong Friendship](https://leetcode.com/problems/strong-friendship/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2), Space: O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1949: Strong Friendship
// https://leetcode.com/problems/strong-friendship/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// friendships: [user1_id, user2_id]
	friendships := [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(StrongFriendship(friendships))
}

// Time: O(n^2), Space: O(n^2)
func StrongFriendship(friendships [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	friendSet := make(map[int]map[int]bool)
	for _, f := range friendships {
		a, b := f[0], f[1]
		if friendSet[a] == nil {
			friendSet[a] = make(map[int]bool)
		}
		if friendSet[b] == nil {
			friendSet[b] = make(map[int]bool)
		}
		friendSet[a][b] = true
		friendSet[b][a] = true
	}

	count := 0
	// For each pair of users, check if they have at least 3 common friends
  // Membuat slice untuk menyimpan hasil
	users := make([]int, 0, len(friendSet))
	for u := range friendSet {
		users = append(users, u)
	}

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			if friendSet[a][b] { // they are friends
				continue
			}
			common := 0
			for f := range friendSet[a] {
				if friendSet[b][f] {
					common++
				}
			}
			if common >= 3 {
				count++
			}
		}
	}
	return count
}
```
