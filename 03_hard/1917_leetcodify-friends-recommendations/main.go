package main

// LeetCode #1917: Leetcodify Friends Recommendations
// https://leetcode.com/problems/leetcodify-friends-recommendations/
// Difficulty: Hard [Paid]
//
// Given tables:
// - listens (user_id, song_id, day)
// - friendship (user1_id, user2_id)
//
// Recommend friend pairs (user1_id, user2_id) who are not already friends
// but listened to the same song(s) on the same day >= 3 times.
// Return distinct pairs with user1_id < user2_id.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple case
	listens := [][]int{
		{1, 101, 1},
		{1, 101, 1}, // same day same song same user (but should not double count)
		{1, 101, 1},
		{2, 101, 1},
		{2, 101, 1},
		{2, 101, 1},
		{3, 101, 1},
	}
	friendships := [][]int{
		{1, 3},
	}
	// Users 1 and 2 both listened to song 101 on day 1 >= 3 times, recommend (1,2)
	// Users 2 and 3: user 3 listened to 101 on day 1 only once
	fmt.Println(leetcodifyFriendsRecommendations(listens, friendships))

	// Test case 2: multiple recommendations
	listens = [][]int{
		{1, 101, 1}, {1, 101, 1}, {1, 101, 1},
		{2, 101, 1}, {2, 101, 1}, {2, 101, 1},
		{1, 102, 2}, {1, 102, 2}, {1, 102, 2},
		{3, 102, 2}, {3, 102, 2}, {3, 102, 2},
	}
	friendships = [][]int{
		{1, 2},
	}
	// 1-2 already friends, skip
	// 1-3 both listened to 102 on day 2 >= 3 times, recommend (1,3)
	fmt.Println(leetcodifyFriendsRecommendations(listens, friendships))

	// Test case 3: no recommendations
	fmt.Println(leetcodifyFriendsRecommendations([][]int{{1, 101, 1}, {2, 102, 1}}, [][]int{}))

	// Test case 4: empty tables
	fmt.Println(leetcodifyFriendsRecommendations([][]int{}, [][]int{}))
}

// leetcodifyFriendsRecommendations returns [][]int{{user1_id, user2_id}, ...}
func leetcodifyFriendsRecommendations(listens [][]int, friendships [][]int) [][]int {
	// Count listens per user per song per day
	type key struct {
		user, song, day int
	}
	count := make(map[key]int)
	for _, l := range listens {
		k := key{user: l[0], song: l[1], day: l[2]}
		count[k]++
	}

	// Build existing friendships set
	existingFriends := make(map[[2]int]bool)
	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if u1 > u2 {
			u1, u2 = u2, u1
		}
		existingFriends[[2]int{u1, u2}] = true
	}

	// For each (song, day) pair, find all users with >= 3 listens
	songDayUsers := make(map[[2]int]map[int]bool) // (song, day) -> set of users with >= 3 listens
	for k, c := range count {
		if c >= 3 {
			sd := [2]int{k.song, k.day}
			if songDayUsers[sd] == nil {
				songDayUsers[sd] = make(map[int]bool)
			}
			songDayUsers[sd][k.user] = true
		}
	}

	// Generate recommendations
	recSet := make(map[[2]int]bool)
	for _, users := range songDayUsers {
		// All pairs of users who listened to this (song, day) >= 3 times
		var userList []int
		for u := range users {
			userList = append(userList, u)
		}
		sort.Ints(userList)
		for i := 0; i < len(userList); i++ {
			for j := i + 1; j < len(userList); j++ {
				u1, u2 := userList[i], userList[j]
				pair := [2]int{u1, u2}
				if !existingFriends[pair] {
					recSet[pair] = true
				}
			}
		}
	}

	// Build sorted result
	var result [][]int
	for p := range recSet {
		result = append(result, []int{p[0], p[1]})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})

	if result == nil {
		return [][]int{}
	}
	return result
}
