# 1892 — Page Recommendations Ii

## Deskripsi

**Soal:** [1892. Page Recommendations Ii](https://leetcode.com/problems/page-recommendations-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1892: Page Recommendations II
// https://leetcode.com/problems/page-recommendations-ii/
// Difficulty: Hard [Paid]
//
// Given tables: Friendship (user1_id, user2_id), Likes (user_id, page_id),
// and Users (user_id), for each user recommend pages that:
// - Are liked by friends of the user's friends (2nd degree friends)
// - ARE NOT already liked by the user
// - ARE NOT already liked by the user's direct friends
// Order by user_id, page_id.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: basic
	friendships := [][]int{
		{1, 2}, {1, 3}, {1, 4}, {2, 3},
	}
	likes := [][]int{
		{1, 101}, {2, 102}, {3, 103}, {4, 104},
		{2, 105},
	}
	users := []int{1}
	// Friends of user 1: {2, 3, 4}
	// Friends of friends (excluding 1 and {2,3,4}): none directly, so empty
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 2:
	friendships = [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	likes = [][]int{
		{1, 101}, {1, 102},
		{2, 102},
		{3, 103},
		{4, 104},
	}
	users = []int{1}
	// Direct friends: {2} (likes 102)
	// Friends of friends: {3} (likes 103) -> recommend 103 (not liked by 1 or direct friends)
	// {4} is friend of 3, also friend-of-friend of 1 -> likes 104 -> recommend 104
	// Result: should get 103, 104
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 3: multiple users
	friendships = [][]int{
		{1, 2}, {2, 3},
	}
	likes = [][]int{
		{1, 101}, {2, 101}, {3, 102},
	}
	users = []int{1, 2}
	fmt.Println(pageRecommendationsIi(friendships, likes, users))

	// Test case 4: empty
	fmt.Println(pageRecommendationsIi([][]int{}, [][]int{}, []int{1}))
}

// pageRecommendationsIi returns [][]int{{user_id, page_id}, ...}
func pageRecommendationsIi(friendships [][]int, likes [][]int, users []int) [][]int {
	// Build friendship graph
  // Membuat map untuk pencarian O(1): key → value
	friends := make(map[int]map[int]bool) // user -> set of friends
	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if friends[u1] == nil {
			friends[u1] = make(map[int]bool)
		}
		if friends[u2] == nil {
			friends[u2] = make(map[int]bool)
		}
		friends[u1][u2] = true
		friends[u2][u1] = true
	}

	// Build likes mapping
  // Membuat map untuk pencarian O(1): key → value
	userLikes := make(map[int]map[int]bool) // user -> set of liked pages
	for _, l := range likes {
		uid, pid := l[0], l[1]
		if userLikes[uid] == nil {
			userLikes[uid] = make(map[int]bool)
		}
		userLikes[uid][pid] = true
	}

	var result [][]int

	for _, user := range users {
		// Find direct friends
		directFriends := friends[user]
		if directFriends == nil {
			directFriends = make(map[int]bool)
		}

		// Find friend-of-friend pages to recommend
  // Membuat map untuk pencarian O(1): key → value
		recommended := make(map[int]bool)

		// For each friend of the user
		for friend := range directFriends {
			// For each friend of that friend
			for fof := range friends[friend] {
				if fof == user || directFriends[fof] {
					continue // skip user and direct friends
				}
				// Add pages liked by fof
				for page := range userLikes[fof] {
					recommended[page] = true
				}
			}
		}

		// Remove pages already liked by user
		for page := range userLikes[user] {
			delete(recommended, page)
		}

		// Remove pages liked by direct friends
		for friend := range directFriends {
			for page := range userLikes[friend] {
				delete(recommended, page)
			}
		}

		// Build sorted result
		var pages []int
		for p := range recommended {
			pages = append(pages, p)
		}
		sort.Ints(pages)
		for _, p := range pages {
			result = append(result, []int{user, p})
		}
	}

	if result == nil {
		return [][]int{}
	}
	return result
}
```
