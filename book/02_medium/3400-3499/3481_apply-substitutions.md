# 3481 — Apply Substitutions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ApplySubstitutions(s string, subs map[string]string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3481: Apply Substitutions
// https://leetcode.com/problems/apply-substitutions/
// Difficulty: Medium [Paid]
// Complexity: O(n * m) time, O(n) space

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", ApplySubstitutions("Hello %name%", map[string]string{"name": "World"}))
	// Test case 2
	fmt.Println("Test 2:", ApplySubstitutions("%a%%b%", map[string]string{"a": "foo", "b": "bar"}))
	// Test case 3
	fmt.Println("Test 3:", ApplySubstitutions("No placeholders", map[string]string{}))
}

func ApplySubstitutions(s string, subs map[string]string) string {
	result := s
	for key, val := range subs {
		result = strings.ReplaceAll(result, "%"+key+"%", val)
	}
	return result
}
```
