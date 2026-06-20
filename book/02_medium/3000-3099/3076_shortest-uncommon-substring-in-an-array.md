# 3076 — Shortest Uncommon Substring In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestUncommonSubstring(arr []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L^2)  
**Kompleksitas Ruang:** O(n * L^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3076: Shortest Uncommon Substring in an Array
// https://leetcode.com/problems/shortest-uncommon-substring-in-an-array/
// Difficulty: Medium
// Time: O(n * L^2) | Space: O(n * L^2)

import "fmt"

func main() {
	fmt.Println(shortestUncommonSubstring([]string{"cab", "ad", "bad", "c"}))
	fmt.Println(shortestUncommonSubstring([]string{"abc", "bcd", "abcd"}))
}

func shortestUncommonSubstring(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)

	for i := 0; i < n; i++ {
		subs := map[string]bool{}
		for l := 0; l < len(arr[i]); l++ {
			for r := l + 1; r <= len(arr[i]); r++ {
				subs[arr[i][l:r]] = true
			}
		}
		best := ""
		for s := range subs {
			common := false
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if contains(arr[j], s) {
					common = true
					break
				}
			}
			if !common {
				if best == "" || len(s) < len(best) || (len(s) == len(best) && s < best) {
					best = s
				}
			}
		}
		ans[i] = best
	}
	return ans
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
```
