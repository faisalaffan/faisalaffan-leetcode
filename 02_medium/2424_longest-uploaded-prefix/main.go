package main

// LeetCode #2424: Longest Uploaded Prefix
// https://leetcode.com/problems/longest-uploaded-prefix/
// Difficulty: Medium
// Time: O(1) amortized | Space: O(n)
// Track uploaded videos. Longest prefix = longest 1..k where all uploaded.

import "fmt"

type LUPrefix struct {
	uploaded []bool
	longest  int
}

func main() {
	lu := Constructor(4)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(3)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(1)
	fmt.Println(lu.Longest()) // 1
	lu.Upload(2)
	fmt.Println(lu.Longest()) // 3
}

func Constructor(n int) LUPrefix {
	return LUPrefix{uploaded: make([]bool, n+2)}
}

func (l *LUPrefix) Upload(video int) {
	l.uploaded[video] = true
	for l.uploaded[l.longest+1] {
		l.longest++
	}
}

func (l *LUPrefix) Longest() int {
	return l.longest
}
