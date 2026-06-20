# 3045 — Count Prefix And Suffix Pairs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPrefixSuffixPairs(words []string) int64
```

> **💡 Hint:** Trie with paired characters

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3045: Count Prefix and Suffix Pairs II
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/
// Difficulty: Hard
//
// Approach: Trie with paired characters
// For each word, simultaneously traverse prefix char (s[i]) and suffix char (s[n-1-i]).
// The trie stores pairs [prefixChar, suffixChar]. At each node, cnt tracks how many
// words have this prefix-suffix pair. For each word, we sum cnt at each matched node
// (these are previous words that match both prefix and suffix).

import "fmt"

type trieNode3045 struct {
	son map[[2]byte]*trieNode3045
	cnt int
}

func countPrefixSuffixPairs(words []string) int64 {
	var ans int64
	root := &trieNode3045{son: make(map[[2]byte]*trieNode3045)}
	for _, s := range words {
		cur := root
		n := len(s)
		for i := 0; i < n; i++ {
			key := [2]byte{s[i], s[n-1-i]}
			if cur.son[key] == nil {
				cur.son[key] = &trieNode3045{son: make(map[[2]byte]*trieNode3045)}
			}
			cur = cur.son[key]
			ans += int64(cur.cnt)
		}
		cur.cnt++
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println("Example 1:", countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
	// Expected: 4

	// Example 2
	fmt.Println("Example 2:", countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"}))
	// Expected: 2

	// No matches
	fmt.Println("No matches:", countPrefixSuffixPairs([]string{"a", "b", "c"}))
	// Expected: 0

	// All same single char
	fmt.Println("All same:", countPrefixSuffixPairs([]string{"a", "a", "a"}))
	// Expected: 3 (3 pairs: 0-1, 0-2, 1-2)

	// Single word
	fmt.Println("Single:", countPrefixSuffixPairs([]string{"hello"}))
	// Expected: 0

	// Longer words
	fmt.Println("Longer:", countPrefixSuffixPairs([]string{"abc", "abcabc", "abcabcabc"}))
	// Expected depends on prefix-suffix matching

	// Verify pair counting
	fmt.Println("Verify:", countPrefixSuffixPairs([]string{"ab", "ab"}))
	// "ab" prefix "a", suffix "b" → pair (a,b)
	// "ab" prefix "ab", suffix "ab" → pair (a,a), (b,b)
}
```
