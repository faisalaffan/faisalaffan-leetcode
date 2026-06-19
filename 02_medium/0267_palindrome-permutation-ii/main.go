package main

// LeetCode #267: Palindrome Permutation II
// https://leetcode.com/problems/palindrome-permutation-ii/
// Difficulty: Medium [Paid]
// Time: O((n/2)!), Space: O(n)

import "fmt"

func generatePalindromes(s string) []string {
	charCount := make([]byte, 128)
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	oddChar := byte(0)
	oddCount := 0
	for i := range charCount {
		if charCount[i]%2 == 1 {
			oddChar = byte(i)
			oddCount++
		}
	}

	if oddCount > 1 {
		return nil
	}

	half := []byte{}
	for i := range charCount {
		for j := 0; j < int(charCount[i])/2; j++ {
			half = append(half, byte(i))
		}
	}

	result := []string{}
	used := make([]bool, len(half))
	var backtrack func(path []byte)
	backtrack = func(path []byte) {
		if len(path) == len(half) {
			pal := string(path)
			rev := ""
			for i := len(path) - 1; i >= 0; i-- {
				rev += string(path[i])
			}
			if oddChar != 0 {
				pal += string(oddChar)
			}
			result = append(result, pal+rev)
			return
		}

		for i := 0; i < len(half); i++ {
			if used[i] {
				continue
			}
			if i > 0 && half[i] == half[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, half[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]byte{})
	return result
}

func main() {
	fmt.Println(generatePalindromes("aabb"))
	fmt.Println(generatePalindromes("abc"))
	fmt.Println(generatePalindromes("a"))
}
