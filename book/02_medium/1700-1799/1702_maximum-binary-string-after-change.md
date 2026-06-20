# 1702 — Maximum Binary String After Change

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumBinaryString(binary string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1702: Maximum Binary String After Change
// https://leetcode.com/problems/maximum-binary-string-after-change/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumBinaryString(binary string) string {
	n := len(binary)
	zeros := 0
	firstZero := -1

	for i, ch := range binary {
		if ch == '0' {
			zeros++
			if firstZero == -1 {
				firstZero = i
			}
		}
	}

	if zeros <= 1 {
		return binary
	}

	// Result: all 1s except position (firstZero + zeros - 1)
	result := make([]byte, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		if i == firstZero+zeros-1 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(maximumBinaryString("000110")) // Expected: "111011"
	fmt.Println(maximumBinaryString("01"))     // Expected: "01"
	fmt.Println(maximumBinaryString("10"))     // Expected: "10"
}
```
