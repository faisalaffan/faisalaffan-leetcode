# 3088 — Make String Anti Palindrome

## Deskripsi

**Soal:** [3088. Make String Anti Palindrome](https://leetcode.com/problems/make-string-anti-palindrome/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func makeStringAntiPalindrome(s string) int`

> **Ide Kunci:** Minimum character changes to make a string anti-palindrome.

## Solusi Go

```go
package main

// LeetCode #3088: Make String Anti-Palindrome
// https://leetcode.com/problems/make-string-anti-palindrome/
// Difficulty: Hard
//
// Approach: Minimum character changes to make a string anti-palindrome.
// An anti-palindrome satisfies s[i] != s[n-1-i] for all i < n/2.
// For each symmetric pair, if characters are equal, we must change one of them.
// For any pair with equal chars, we can always change one to a different character
// since there are 26 lowercase letters. However, for n = 1, the single character
// pairs with itself, making it impossible to satisfy s[0] != s[0], return -1.
//
// Edge case: if n is odd, the middle character (at index n/2) has no pair
// (since n-1-i == i), so it's excluded from the condition.

import "fmt"

func makeStringAntiPalindrome(s string) int {
	n := len(s)

	// Single character: impossible since s[0] == reverse(s[0]) always
	if n <= 1 {
		return -1
	}

	changes := 0
	for i := 0; i < n/2; i++ {
		if s[i] == s[n-1-i] {
			changes++
		}
	}
	return changes
}

func main() {
	// Example 1: "ab" -> already anti-palindrome (a != b)
	fmt.Println("Test 1 (ab):", makeStringAntiPalindrome("ab"))
	// Expected: 0

	// Example 2: "aa" -> need to change one char
	fmt.Println("Test 2 (aa):", makeStringAntiPalindrome("aa"))
	// Expected: 1

	// Example 3: "aba" -> middle char free, need to change one of pair a==a
	fmt.Println("Test 3 (aba):", makeStringAntiPalindrome("aba"))
	// Expected: 1

	// Example 4: "a" -> single char impossible
	fmt.Println("Test 4 (a):", makeStringAntiPalindrome("a"))
	// Expected: -1

	// Example 5: "abc" -> all pairs different (a!=c)
	fmt.Println("Test 5 (abc):", makeStringAntiPalindrome("abc"))
	// Expected: 0

	// Example 6: "aaaa" -> two pairs, both equal
	fmt.Println("Test 6 (aaaa):", makeStringAntiPalindrome("aaaa"))
	// Expected: 2

	// Example 7: "abca" -> pair (0,3): a==a, pair (1,2): b!=c
	fmt.Println("Test 7 (abca):", makeStringAntiPalindrome("abca"))
	// Expected: 1

	// Example 8: "racecar" -> palindrome, 3 pairs
	// (0,6): r==r, (1,5): a==a, (2,4): c==c
	fmt.Println("Test 8 (racecar):", makeStringAntiPalindrome("racecar"))
	// Expected: 3

	// Example 9: "" -> empty string
	fmt.Println("Test 9 (empty):", makeStringAntiPalindrome(""))
	// Expected: -1

	// Example 10: "xyz" -> pairs: x!=z
	fmt.Println("Test 10 (xyz):", makeStringAntiPalindrome("xyz"))
	// Expected: 0

	// Example 11: "aabaa" -> palindrome, pairs: (0,4): a==a, (1,3): a==a
	fmt.Println("Test 11 (aabaa):", makeStringAntiPalindrome("aabaa"))
	// Expected: 2

	// Example 12: "abbc" -> pairs: (0,3): a!=c, (1,2): b==b
	fmt.Println("Test 12 (abbc):", makeStringAntiPalindrome("abbc"))
	// Expected: 1

	// Example 13: "zz" -> change one
	fmt.Println("Test 13 (zz):", makeStringAntiPalindrome("zz"))
	// Expected: 1
}
```
