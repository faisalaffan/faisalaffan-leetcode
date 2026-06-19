package main

// LeetCode #721: Accounts Merge
// https://leetcode.com/problems/accounts-merge/
// Difficulty: Medium
// Time: O(nk * alpha(nk))
// Space: O(nk)

import (
	"fmt"
	"sort"
)

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	owner := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]
		for _, email := range acc[1:] {
			parent[email] = email
			owner[email] = name
			union(firstEmail, email)
		}
	}

	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}

	result := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		merged := append([]string{owner[root]}, emails...)
		result = append(result, merged)
	}

	return result
}
