# 2067 — Number Of Equal Count Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func equalCountSubstrings(s string, count int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n * alphabet)  
**Kompleksitas Ruang:** O(alphabet)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2067: Number of Equal Count Substrings
// https://leetcode.com/problems/number-of-equal-count-substrings/
// Difficulty: Medium [Paid]
// Time: O(n * alphabet) | Space: O(alphabet)

import "fmt"

func equalCountSubstrings(s string, count int) int {
	result := 0
	// Try different numbers of distinct characters
	for distinct := 1; distinct <= 26 && distinct*count <= len(s); distinct++ {
  // Alokasi slice integer
		freq := make([]int, 26)
		unique := 0
		exactCount := 0

  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s); i++ {
			idx := int(s[i] - 'a')
			if freq[idx] == 0 {
				unique++
			}
			freq[idx]++
			if freq[idx] == count {
				exactCount++
			}

			// Remove leftmost when window too big
			if i >= distinct*count {
				left := int(s[i-distinct*count] - 'a')
				if freq[left] == count {
					exactCount--
				}
				freq[left]--
				if freq[left] == 0 {
					unique--
				}
			}

			if unique == distinct && exactCount == distinct {
				result++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalCountSubstrings("aaabc", 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", equalCountSubstrings("abcd", 1))
	// Expected: 10

	// Test case 3
	fmt.Println("Test 3:", equalCountSubstrings("aabbcc", 2))
	// Expected: 6
}
```
