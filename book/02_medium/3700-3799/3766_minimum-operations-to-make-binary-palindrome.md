# 3766 — Minimum Operations To Make Binary Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func init() 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * log M)  
**Kompleksitas Ruang:** O(M)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3766: Minimum Operations to Make Binary Palindrome
// https://leetcode.com/problems/minimum-operations-to-make-binary-palindrome/
// Difficulty: Medium
// Time: O(n * log M) | Space: O(M)

import (
	"fmt"
	"sort"
	"strconv"
)

var binaryPalindromes []int

func init() {
	for i := 0; i < (1 << 14); i++ {
		s := strconv.FormatInt(int64(i), 2)
		if isPalindromeStr(s) {
			binaryPalindromes = append(binaryPalindromes, i)
		}
	}
}

func isPalindromeStr(s string) bool {
	l, r := 0, len(s)-1
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func minimumOperationsToMakeBinaryPalindrome(nums []int) []int {
  // Alokasi slice integer
	ans := make([]int, len(nums))
	for idx, x := range nums {
		pos := sort.SearchInts(binaryPalindromes, x)
		best := 1 << 30
		if pos < len(binaryPalindromes) {
			if binaryPalindromes[pos]-x < best {
				best = binaryPalindromes[pos] - x
			}
		}
		if pos > 0 {
			if x-binaryPalindromes[pos-1] < best {
				best = x - binaryPalindromes[pos-1]
			}
		}
		ans[idx] = best
	}
	return ans
}

func main() {
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{1, 2, 3, 4, 5}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{10, 20}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{7}))
}
```
