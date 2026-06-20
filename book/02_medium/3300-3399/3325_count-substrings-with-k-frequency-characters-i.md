# 3325 — Count Substrings With K Frequency Characters I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubstrings(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3325: Count Substrings With K-Frequency Characters I
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-i/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2)) // 4
	fmt.Println(numberOfSubstrings("abcde", 1)) // 15
	fmt.Println(numberOfSubstrings("aaaa", 2))  // 3
}

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	cnt := [26]int{}
	ans, left := 0, 0

	for right := 0; right < n; right++ {
		idx := s[right] - 'a'
		cnt[idx]++

		for cnt[idx] >= k {
			ans += n - right
			cnt[s[left]-'a']--
			left++
		}
	}

	return ans
}
```
