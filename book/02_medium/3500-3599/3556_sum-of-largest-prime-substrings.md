# 3556 — Sum Of Largest Prime Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isPrime(x int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3556: Sum of Largest Prime Substrings
// https://leetcode.com/problems/sum-of-largest-prime-substrings/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import (
	"fmt"
	"strconv"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SumOfLargestPrimeSubstrings("237"))
	// Test case 2
	fmt.Println("Test 2:", SumOfLargestPrimeSubstrings("1234"))
	// Test case 3
	fmt.Println("Test 3:", SumOfLargestPrimeSubstrings("111"))
}

func SumOfLargestPrimeSubstrings(s string) int {
	sum := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s) && j-i <= 6; j++ { // at most 6 digits to avoid overflow
			num, _ := strconv.Atoi(s[i:j])
			if isPrime(num) {
				sum += num
			}
		}
	}
	return sum
}
```
