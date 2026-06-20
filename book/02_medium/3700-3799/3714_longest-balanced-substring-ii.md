# 3714 — Longest Balanced Substring Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestBalancedSubstringIi(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3714: Longest Balanced Substring II
// https://leetcode.com/problems/longest-balanced-substring-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestBalancedSubstringIi(s string) int {
	n := len(s)
	ans := 0

	// Case 1: single char run
	i := 0
	for i < n {
		start := i
		for i < n && s[i] == s[start] {
			i++
		}
		if i-start > ans {
			ans = i - start
		}
	}

	// Case 2: exactly two chars
	solvePair := func(x, y byte) {
		i := 0
		for i < n {
			pos := map[int]int{0: i - 1}
			d := 0
			for i < n && (s[i] == x || s[i] == y) {
				if s[i] == x {
					d++
				} else {
					d--
				}
				if firstIdx, ok := pos[d]; ok {
					if i-firstIdx > ans {
						ans = i - firstIdx
					}
				} else {
					pos[d] = i
				}
				i++
			}
			if i < n {
				i++ // skip third char
			}
		}
	}

	solvePair('a', 'b')
	solvePair('a', 'c')
	solvePair('b', 'c')

	// Case 3: all three chars
  // Membuat map (HashMap) — pencarian O(1)
	pos3 := make(map[[2]int]int)
	pos3[[2]int{0, 0}] = -1
	cnt := [3]int{} // a, b, c
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			cnt[0]++
		case 'b':
			cnt[1]++
		case 'c':
			cnt[2]++
		}
		state := [2]int{cnt[0] - cnt[1], cnt[1] - cnt[2]}
		if firstIdx, ok := pos3[state]; ok {
			if i-firstIdx > ans {
				ans = i - firstIdx
			}
		} else {
			pos3[state] = i
		}
	}

	return ans
}

func main() {
	fmt.Println(longestBalancedSubstringIi("abcabc"))
	fmt.Println(longestBalancedSubstringIi("aabbcc"))
	fmt.Println(longestBalancedSubstringIi("aaa"))
}
```
