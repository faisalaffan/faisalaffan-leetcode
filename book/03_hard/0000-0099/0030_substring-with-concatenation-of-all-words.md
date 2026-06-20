# 0030 — Substring With Concatenation Of All Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findSubstring(s string, words []string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #30: Substring with Concatenation of All Words
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// Difficulty: Hard

import "fmt"

// findSubstring finds all starting indices where concatenation of all words matches.
// All words have the same length.
//
// Complexity: O(n * m) time, O(k) space where n = len(s), m = len(words), k = number of unique words
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalLen := wordLen * len(words)

	if len(s) < totalLen {
		return []int{}
	}

	// Build frequency map for words
  // Membuat map (HashMap) — pencarian O(1)
	wordFreq := make(map[string]int)
	for _, w := range words {
		wordFreq[w]++
	}

	var result []int

	// Slide over the string in wordLen-sized chunks
	for i := 0; i < wordLen; i++ {
		left := i
		right := i
  // Membuat map (HashMap) — pencarian O(1)
		windowFreq := make(map[string]int)
		count := 0

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordFreq[word]; exists {
				windowFreq[word]++
				count++

				// If we have too many of this word, slide left
				for windowFreq[word] > wordFreq[word] {
					leftWord := s[left : left+wordLen]
					windowFreq[leftWord]--
					count--
					left += wordLen
				}

				// If we've matched all words, record the starting index
				if count == len(words) {
					result = append(result, left)
					// Slide left window by one word
					leftWord := s[left : left+wordLen]
					windowFreq[leftWord]--
					count--
					left += wordLen
				}
			} else {
				// Reset window
				windowFreq = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}

func main() {
	// Test case from LeetCode
	fmt.Println("Test 1:", findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0, 9]

	// Additional test cases
	fmt.Println("Test 2:", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
	fmt.Println("Test 3:", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))            // [6, 9, 12]
}
```
