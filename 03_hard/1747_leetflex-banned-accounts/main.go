package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type LogInfo struct {
	AccountId int
	IpAddress int
	Login     int
	Logout    int
}

func main() {
	logs1 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 2, 4},
		{2, 1, 6, 8},
	}
	result1 := leetflexBannedAccounts(logs1)
	fmt.Printf("Test 1 - Banned accounts: %v (Expected: [1])\n\n", result1)

	logs2 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 6, 8},
		{2, 1, 1, 5},
		{2, 2, 6, 8},
		{3, 1, 1, 10},
		{3, 2, 5, 15},
	}
	result2 := leetflexBannedAccounts(logs2)
	fmt.Printf("Test 2 - Banned accounts: %v (Expected: [3])\n\n", result2)

	logs3 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 6, 8},
		{2, 1, 1, 5},
		{2, 2, 6, 8},
	}
	result3 := leetflexBannedAccounts(logs3)
	fmt.Printf("Test 3 - Banned accounts: %v (Expected: [])\n", result3)
}

func leetflexBannedAccounts(logs []LogInfo) []int {
	// Group sessions by account
	byAccount := make(map[int][]LogInfo)
	for _, l := range logs {
		byAccount[l.AccountId] = append(byAccount[l.AccountId], l)
	}

	banned := make([]int, 0)
	for accountId, sessions := range byAccount {
		// Merge overlapping sessions for each IP, then check for overlap across IPs
		ipSessions := make(map[int][]LogInfo)
		for _, s := range sessions {
			ipSessions[s.IpAddress] = append(ipSessions[s.IpAddress], s)
		}

		// Merge per-IP sessions
		merged := make([]LogInfo, 0)
		for _, ipSess := range ipSessions {
			sort.Slice(ipSess, func(i, j int) bool {
				return ipSess[i].Login < ipSess[j].Login
			})
			mergedIP := ipSess[0]
			for i := 1; i < len(ipSess); i++ {
				if ipSess[i].Login <= mergedIP.Logout+1 {
					if ipSess[i].Logout > mergedIP.Logout {
						mergedIP.Logout = ipSess[i].Logout
					}
				} else {
					merged = append(merged, mergedIP)
					mergedIP = ipSess[i]
				}
			}
			merged = append(merged, mergedIP)
		}

		// Sort all merged sessions by login time
		sort.Slice(merged, func(i, j int) bool {
			return merged[i].Login < merged[j].Login
		})

		// Check for overlapping sessions (same account, different IPs)
		for i := 1; i < len(merged); i++ {
			if merged[i].Login <= merged[i-1].Logout {
				banned = append(banned, accountId)
				break
			}
		}
	}

	sort.Ints(banned)
	return banned
}
