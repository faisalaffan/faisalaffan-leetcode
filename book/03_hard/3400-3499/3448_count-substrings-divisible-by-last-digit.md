# 3448 — Count Substrings Divisible By Last Digit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countSubstrings(s string) int64
```

> **💡 Hint:** Iterate through the string, for each position consider it as

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3448: Count Substrings Divisible By Last Digit
// https://leetcode.com/problems/count-substrings-divisible-by-last-digit/
// Difficulty: Hard
//
// Count substrings where the integer formed by the substring is divisible
// by its last digit. The last digit cannot be 0 (division by zero).
//
// Approach: Iterate through the string, for each position consider it as
// the last digit. Check all substrings ending at this position for
// divisibility.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubstrings("12936"))
	// Example 2
	fmt.Println(countSubstrings("5701283"))
	// Example 3: all zeros
	fmt.Println(countSubstrings("1010"))
	// Edge: single digit
	fmt.Println(countSubstrings("5"))
	// Edge: with zeros
	fmt.Println(countSubstrings("0"))
}

func countSubstrings(s string) int64 {
	n := len(s)
	var ans int64

	for j := 0; j < n; j++ {
		lastDigit := int(s[j] - '0')
		if lastDigit == 0 {
			continue
		}
		// Check substrings ending at j
		num := 0
		for i := j; i >= 0; i-- {
			digit := int(s[i] - '0')
			// Build number from left to right
			num = (digit + num*10) % lastDigit
			if num == 0 {
				ans++
			}
		}
	}

	return ans
}
```
