# 3297 — Count Substrings That Can Be Rearranged To Contain A String I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func validSubstringCount(word1 string, word2 string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3297: Count Substrings That Can Be Rearranged to Contain a String I
// https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/
// Difficulty: Medium
// Time: O(n + m) Space: O(1)

import "fmt"

func main() {
	fmt.Println(validSubstringCount("bcca", "abc")) // 1
	fmt.Println(validSubstringCount("abcabc", "abc")) // 10
	fmt.Println(validSubstringCount("a", "aa"))       // 0
}

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	var cnt [26]int
	need := 0
	for _, ch := range word2 {
		idx := ch - 'a'
		if cnt[idx] == 0 {
			need++
		}
		cnt[idx]++
	}

	var win [26]int
	var ans int64
	left := 0

	for _, ch := range word1 {
		idx := int(ch - 'a')
		win[idx]++
		if win[idx] == cnt[idx] {
			need--
		}

		for need == 0 {
			leftIdx := int(word1[left] - 'a')
			if win[leftIdx] == cnt[leftIdx] {
				need++
			}
			win[leftIdx]--
			left++
		}

		ans += int64(left)
	}

	return ans
}
```
