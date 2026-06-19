package main

// LeetCode #752: Open the Lock
// https://leetcode.com/problems/open-the-lock/
// Difficulty: Medium
// Time: O(10^4 * 8) ~ O(1)
// Space: O(10^4)

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
}

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}

	visited := make(map[string]bool)
	queue := []string{"0000"}
	visited["0000"] = true
	steps := 0

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			if curr == target {
				return steps
			}

			for j := 0; j < 4; j++ {
				for _, d := range []int{-1, 1} {
					next := []byte(curr)
					next[j] = byte('0' + (int(next[j]-'0')+d+10)%10)
					s := string(next)
					if !visited[s] && !dead[s] {
						visited[s] = true
						queue = append(queue, s)
					}
				}
			}
		}
		queue = queue[n:]
		steps++
	}

	return -1
}
