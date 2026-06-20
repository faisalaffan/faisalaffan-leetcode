# 3305 — Count Of Substrings Containing Every Vowel And K Consonants I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countOfSubstrings(word string, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3305: Count of Substrings Containing Every Vowel and K Consonants I
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))         // 0
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstrings("aeiou", 0))          // 1
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var ans int64
	for i := 0; i < n; i++ {
  // Membuat map (HashMap) — pencarian O(1)
		vowelCnt := make(map[byte]int)
		cons := 0
		for j := i; j < n; j++ {
			if isVowel(word[j]) {
				vowelCnt[word[j]]++
			} else {
				cons++
				if cons > k {
					break
				}
			}
			if cons == k && len(vowelCnt) == 5 {
				ans++
			}
		}
	}
	return ans
}
```
