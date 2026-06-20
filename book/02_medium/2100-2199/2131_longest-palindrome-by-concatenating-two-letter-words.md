# 2131 — Longest Palindrome By Concatenating Two Letter Words

## Deskripsi

**Soal:** [2131. Longest Palindrome By Concatenating Two Letter Words](https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func longestPalindrome(words []string) int`

## Solusi Go

```go
package main

// LeetCode #2131: Longest Palindrome by Concatenating Two-Letter Words
// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestPalindrome(words []string) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	length := 0
	centerUsed := false

	for w, count := range freq {
		if count == 0 {
			continue
		}
		rev := string([]byte{w[1], w[0]})

		if w == rev {
			// Same letter pair like "aa"
			pairs := count / 2
			length += pairs * 4
			if count%2 == 1 && !centerUsed {
				length += 2
				centerUsed = true
			}
			freq[w] = 0
		} else if revCount, ok := freq[rev]; ok && revCount > 0 {
			pairs := count
			if revCount < pairs {
				pairs = revCount
			}
			length += pairs * 4
			freq[w] -= pairs
			freq[rev] -= pairs
		}
	}

	return length
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestPalindrome([]string{"lc", "cl", "gg"}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", longestPalindrome([]string{"cc", "ll", "xx"}))
	// Expected: 2
}
```
