# 3306 — Count Of Substrings Containing Every Vowel And K Consonants Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countOfSubstringsII(word string, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3306: Count of Substrings Containing Every Vowel and K Consonants II
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstringsII("aeioqq", 1))         // 0
	fmt.Println(countOfSubstringsII("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstringsII("aeiou", 0))          // 1
}

func countOfSubstringsII(word string, k int) int64 {
	return atLeastK(word, k) - atLeastK(word, k+1)
}

func atLeastK(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

  // Membuat map (HashMap) — pencarian O(1)
	vowelCnt := make(map[byte]int)
	cons := 0
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		c := word[right]
		if isVowel(c) {
			vowelCnt[c]++
		} else {
			cons++
		}

		for len(vowelCnt) == 5 && cons >= k {
			out := word[left]
			if isVowel(out) {
				vowelCnt[out]--
				if vowelCnt[out] == 0 {
					delete(vowelCnt, out)
				}
			} else {
				cons--
			}
			left++
		}
		ans += int64(left)
	}
	return ans
}
```
