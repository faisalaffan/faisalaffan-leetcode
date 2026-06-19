package main

// LeetCode #2075: Decode the Slanted Ciphertext
// https://leetcode.com/problems/decode-the-slanted-ciphertext/
// Difficulty: Medium
// Time: O(r*c) | Space: O(r*c)

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []byte(encodedText[i*cols : (i+1)*cols])
	}

	var result strings.Builder
	for startCol := 0; startCol < cols; startCol++ {
		r, c := 0, startCol
		for r < rows && c < cols {
			result.WriteByte(grid[r][c])
			r++
			c++
		}
	}

	// Trim trailing spaces
	s := result.String()
	s = strings.TrimRight(s, " ")
	return s
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", decodeCiphertext("ch   ie   pr", 3))
	// Expected: "cipher"

	// Test case 2
	fmt.Println("Test 2:", decodeCiphertext("iveo    eed   l te   olc", 4))
	// Expected: "i love leetcode"

	// Test case 3
	fmt.Println("Test 3:", decodeCiphertext("coding", 1))
	// Expected: "coding"
}
