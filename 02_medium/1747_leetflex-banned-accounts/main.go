package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Login struct {
	AccountID int
	IPAddress string
	LoginTime int
}

func findBanned(logins []Login) []int {
	// Group by account
	groups := make(map[int][]Login)
	for _, l := range logins {
		groups[l.AccountID] = append(groups[l.AccountID], l)
	}

	banned := make(map[int]bool)
	for accID, records := range groups {
		sort.Slice(records, func(i, j int) bool {
			return records[i].LoginTime < records[j].LoginTime
		})
		// Track latest login time per IP for this account
		lastTime := make(map[string]int)
		for _, r := range records {
			if prevTime, ok := lastTime[r.IPAddress]; ok {
				// Same IP, update last time
				_ = prevTime
			}
			lastTime[r.IPAddress] = r.LoginTime
		}

		// Check if any IP has concurrent sessions
		active := make(map[string]int)
		for _, r := range records {
			if _, ok := active[r.IPAddress]; ok {
				// Check if there's a different IP with an active session
				for ip, t := range active {
					if ip != r.IPAddress && t < r.LoginTime {
						banned[accID] = true
					}
				}
				// End current session for this IP and start new one
				delete(active, r.IPAddress)
			}
			active[r.IPAddress] = r.LoginTime
		}
	}

	result := make([]int, 0, len(banned))
	for id := range banned {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func main() {
	logins := []Login{
		{1, "IP1", 1},
		{1, "IP2", 2},
		{1, "IP1", 3},
	}
	fmt.Println(findBanned(logins))
}
