package main

// LeetCode #3335: Total Characters in String After Transformations I
// https://leetcode.com/problems/total-characters-in-string-after-transformations-i/
// Difficulty: Medium
// Time: O(n + 26t) Space: O(26)

import "fmt"

func main() {
	fmt.Println(lengthAfterTransformations("ab", 1)) // 2
	fmt.Println(lengthAfterTransformations("z", 1))  // 2
	fmt.Println(lengthAfterTransformations("az", 2)) // 5
}

func lengthAfterTransformations(s string, t int) int {
	const mod = 1_000_000_007
	freq := [26]int64{}
	for _, c := range s {
		freq[c-'a']++
	}

	for ; t > 0; t-- {
		next := [26]int64{}
		for i, cnt := range freq {
			if cnt == 0 {
				continue
			}
			if i == 25 { // 'z' -> "ab"
				next[0] = (next[0] + cnt) % mod
				next[1] = (next[1] + cnt) % mod
			} else {
				next[i+1] = (next[i+1] + cnt) % mod
			}
		}
		freq = next
	}

	var ans int64
	for _, cnt := range freq {
		ans = (ans + cnt) % mod
	}
	return int(ans)
}
