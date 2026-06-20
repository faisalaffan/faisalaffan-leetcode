# 2156 — Find Substring With Given Hash Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func subStrHash(s string, power int, mod int, k int, hashValue int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2156: Find Substring With Given Hash Value
// https://leetcode.com/problems/find-substring-with-given-hash-value/
// Difficulty: Hard
//
// Rolling hash from left to right. Precompute power^(k-1) % mod for sliding.
// h(i) = s[i]*p^{k-1} + s[i+1]*p^{k-2} + ... + s[i+k-1]*p^0
// h(i+1) = (h(i) - s[i]*p^{k-1}) * power + s[i+k]

import "fmt"

func main() {
	fmt.Println(subStrHash("leetcode", 7, 20, 2, 0))  // "ee"
	fmt.Println(subStrHash("fbxzaad", 31, 100, 3, 32)) // "" (no match)
	fmt.Println(subStrHash("xqgcas", 4, 7, 3, 4))     // "xqg"
	fmt.Println(subStrHash("helloworld", 10, 1000, 3, 862)) // "hel"
}

func subStrHash(s string, power int, mod int, k int, hashValue int) string {
	n := len(s)

	// pk1 = power^(k-1) % mod
	pk1 := 1
	for i := 0; i < k-1; i++ {
		pk1 = (pk1 * power) % mod
	}
	// pk = power^k % mod (not strictly needed, pk1 * power)

	// Compute hash of first window
	cur := 0
	for i := 0; i < k; i++ {
		cur = (cur*power + int(s[i]-'a'+1)) % mod
	}
	if cur == hashValue {
		return s[:k]
	}

	// Slide window left to right
	for i := 1; i <= n-k; i++ {
		// h(i) = (h(i-1) - s[i-1]*p^{k-1}) * power + s[i+k-1]
		cur = (cur - (int(s[i-1]-'a'+1))*pk1%mod + mod) % mod
		cur = (cur*power + int(s[i+k-1]-'a'+1)) % mod
		if cur == hashValue {
			return s[i : i+k]
		}
	}
	return ""
}

func FindSubstringWithGivenHashValue() any {
	return subStrHash("leetcode", 7, 20, 2, 0)
}
```
