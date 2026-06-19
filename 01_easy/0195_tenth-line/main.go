package main

// LeetCode #195: Tenth Line
// https://leetcode.com/problems/tenth-line/
// Difficulty: Easy

import "fmt"

func TenthLine() string {
	return "sed -n '10p' file.txt"
}

func main() {
	fmt.Println(TenthLine())
}
