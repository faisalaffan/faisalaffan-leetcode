package main

// LeetCode #3889: Mirror Frequency Distance
// https://leetcode.com/problems/mirror-frequency-distance/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count character frequencies. For each unique char, compute mirror
// char and sum absolute frequency differences, counting each pair once.

import "fmt"

func MirrorFrequencyDistance(s string) int {
	freq := make([]int, 36) // 0-25: letters, 26-35: digits
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			freq[ch-'a']++
		} else {
			freq[26+int(ch-'0')]++
		}
	}

	visited := make([]bool, 36)
	ans := 0
	for i := 0; i < 36; i++ {
		if visited[i] || freq[i] == 0 {
			continue
		}
		// Compute mirror
		var mirror int
		if i < 26 {
			mirror = 25 - i // 'a'->'z', 'b'->'y', etc.
		} else {
			mirror = 26 + (9 - (i - 26)) // '0'->'9', '1'->'8', etc.
		}
		visited[i] = true
		if mirror != i {
			visited[mirror] = true
		}
		diff := freq[i] - freq[mirror]
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MirrorFrequencyDistance("ab1z9")) // Expected: 3

	// Example 2
	fmt.Println(MirrorFrequencyDistance("4m7n")) // Expected: 2

	// Example 3
	fmt.Println(MirrorFrequencyDistance("byby")) // Expected: 0
}
