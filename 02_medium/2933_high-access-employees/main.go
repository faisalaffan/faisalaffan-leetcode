package main

// LeetCode #2933: High-Access Employees
// https://leetcode.com/problems/high-access-employees/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findHighAccessEmployees([][]string{{"a","0549"},{"b","0456"},{"a","1234"},{"a","1440"}}))
	fmt.Println(findHighAccessEmployees([][]string{{"d","0002"},{"d","0003"},{"d","0004"},{"d","0005"}}))
}

func findHighAccessEmployees(accessTimes [][]string) (ans []string) {
	d := map[string][]int{}
	for _, e := range accessTimes {
		name, s := e[0], e[1]
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[2:])
		t := h*60 + m
		d[name] = append(d[name], t)
	}
	for name, ts := range d {
		sort.Ints(ts)
		for i := 2; i < len(ts); i++ {
			if ts[i]-ts[i-2] < 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	return
}
