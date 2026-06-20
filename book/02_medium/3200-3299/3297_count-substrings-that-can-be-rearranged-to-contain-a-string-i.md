# 3297 — Count Substrings That Can Be Rearranged To Contain A String I

## Deskripsi

**Soal:** [3297. Count Substrings That Can Be Rearranged To Contain A String I](https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3297: Count Substrings That Can Be Rearranged to Contain a String I
// https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/
// Difficulty: Medium
// Time: O(n + m) Space: O(1)

import "fmt"

func main() {
	fmt.Println(validSubstringCount("bcca", "abc")) // 1
	fmt.Println(validSubstringCount("abcabc", "abc")) // 10
	fmt.Println(validSubstringCount("a", "aa"))       // 0
}

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	var cnt [26]int
	need := 0
	for _, ch := range word2 {
		idx := ch - 'a'
		if cnt[idx] == 0 {
			need++
		}
		cnt[idx]++
	}

	var win [26]int
	var ans int64
	left := 0

	for _, ch := range word1 {
		idx := int(ch - 'a')
		win[idx]++
		if win[idx] == cnt[idx] {
			need--
		}

		for need == 0 {
			leftIdx := int(word1[left] - 'a')
			if win[leftIdx] == cnt[leftIdx] {
				need++
			}
			win[leftIdx]--
			left++
		}

		ans += int64(left)
	}

	return ans
}
```
