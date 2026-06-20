# 2301 — Match Substring After Replacement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func matchReplacement(s string, sub string, mappings [][]byte) bool
```

> **💡 Hint:** Build a directed graph from mappings, compute transitive closure

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Floyd-Warshall

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Floyd-Warshall** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2301: Match Substring After Replacement
// https://leetcode.com/problems/match-substring-after-replacement/
// Difficulty: Hard
//
// Approach: Build a directed graph from mappings, compute transitive closure
// (a->b and b->c implies a->c). Then check if each character in s can be
// transformed to the corresponding character in sub, character by character.

import "fmt"

func main() {
	// Example 1: "fool3e7bar","leet",[["e","3"],["t","7"],["t","8"]] => true
	fmt.Println(matchReplacement("fool3e7bar", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}))
	// Example 2: "fooleetbar","f00l",[["o","0"]] => false
	fmt.Println(matchReplacement("fooleetbar", "f00l", [][]byte{{'o', '0'}}))
	// Example 3: "Fool33tbaR","leet",[["e","3"],["t","7"],["t","8"],["e","E"],["e","e"]] => true
	fmt.Println(matchReplacement("Fool33tbaR", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}, {'e', 'E'}, {'e', 'e'}}))
	// Edge: sub == s
	fmt.Println(matchReplacement("abc", "abc", [][]byte{}))
	// Edge: single char
	fmt.Println(matchReplacement("a", "b", [][]byte{{'a', 'b'}}))
	fmt.Println(matchReplacement("a", "b", [][]byte{}))
}

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	// Build transitive closure: can[a][b] means a can be replaced by b
  // Membuat matriks/slice 2D untuk DP
	can := make([][]bool, 256)
  // Range loop: iterasi dengan indeks + nilai
	for i := range can {
		can[i] = make([]bool, 256)
		can[i][i] = true
	}
	for _, m := range mappings {
		can[m[0]][m[1]] = true
	}

	// Floyd-Warshall for transitive closure
	for k := 0; k < 256; k++ {
		for i := 0; i < 256; i++ {
			if !can[i][k] {
				continue
			}
			for j := 0; j < 256; j++ {
				if can[k][j] {
					can[i][j] = true
				}
			}
		}
	}

	// Slide sub through s
	for start := 0; start <= len(s)-len(sub); start++ {
		match := true
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(sub); i++ {
			if !can[sub[i]][s[start+i]] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
```
