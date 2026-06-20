# 1100 — Find K Length Substrings With No Repeated Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numKLenSubstrNoRepeats(s string, k int) int
```

> **💡 Hint:** Sliding window with frequency array

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1100: Find K-Length Substrings With No Repeated Characters
// https://leetcode.com/problems/find-k-length-substrings-with-no-repeated-characters/
// Difficulty: Medium
//
// Approach: Sliding window with frequency array
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode", 5)) // 6
	fmt.Println(numKLenSubstrNoRepeats("home", 5))              // 0
}

func numKLenSubstrNoRepeats(s string, k int) int {
	if k > len(s) || k > 26 {
		return 0
	}

  // Alokasi slice integer
	freq := make([]int, 26)
	duplicates := 0
	result := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		if freq[s[i]-'a'] == 2 {
			duplicates++
		}

		if i >= k {
			freq[s[i-k]-'a']--
			if freq[s[i-k]-'a'] == 1 {
				duplicates--
			}
		}

		if i >= k-1 && duplicates == 0 {
			result++
		}
	}

	return result
}
```
