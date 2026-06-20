# 2168 — Unique Substrings With Equal Digit Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func equalDigitFrequency(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2168: Unique Substrings With Equal Digit Frequency
// https://leetcode.com/problems/unique-substrings-with-equal-digit-frequency/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func equalDigitFrequency(s string) int {
	n := len(s)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
  // Alokasi slice integer
		freq := make([]int, 10)
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			d := int(s[j] - '0')
			if freq[d] == 0 {
				distinct++
			}
			freq[d]++
			if freq[d] > maxFreq {
				maxFreq = freq[d]
			}
			// All digits appear same frequency iff distinct * maxFreq == totalLen
			if distinct*maxFreq == j-i+1 {
				seen[s[i:j+1]] = true
			}
		}
	}

	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalDigitFrequency("1212"))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", equalDigitFrequency("12321"))
	// Expected: 9
}
```
