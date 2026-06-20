# 1980 — Find Unique Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindDifferentBinaryString(nums []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1980: Find Unique Binary String
// https://leetcode.com/problems/find-unique-binary-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindDifferentBinaryString([]string{"01", "10"}))
	fmt.Println(FindDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(FindDifferentBinaryString([]string{"111", "011", "001"}))
}

// Time: O(n^2), Space: O(n)
func FindDifferentBinaryString(nums []string) string {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[string]bool)
	for _, s := range nums {
		set[s] = true
	}

	// Generate candidates using Cantor diagonal argument
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = '0'
	}
	for {
		s := string(result)
		if !set[s] {
			return s
		}
		// Increment binary string
		j := n - 1
		for j >= 0 && result[j] == '1' {
			result[j] = '0'
			j--
		}
		if j < 0 {
			break
		}
		result[j] = '1'
	}
	return ""
}
```
