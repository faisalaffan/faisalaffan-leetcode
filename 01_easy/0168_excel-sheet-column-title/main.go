package main

// LeetCode #168: Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(log n)
func ConvertToTitle(columnNumber int) string {
	res := make([]byte, 0, 8)
	for columnNumber > 0 {
		columnNumber--
		res = append(res, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(ConvertToTitle(1))
	fmt.Println(ConvertToTitle(28))
	fmt.Println(ConvertToTitle(701))
}
