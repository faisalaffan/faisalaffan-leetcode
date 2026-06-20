# 3955 — Valid Binary Strings With Cost Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidBinaryStringsWithCostLimit(n int, k int) []string
```

> **💡 Hint:** Backtracking. Generate all strings without consecutive 1s,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Backtracking

**Kompleksitas Waktu:** O(2^N)  
**Kompleksitas Ruang:** O(N * 2^N)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3955: Valid Binary Strings With Cost Limit
// https://leetcode.com/problems/valid-binary-strings-with-cost-limit/
// Difficulty: Medium
// Time: O(2^N) | Space: O(N * 2^N)
// Approach: Backtracking. Generate all strings without consecutive 1s,
// filter by cost (sum of indices where s[i]=='1') <= k.

import "fmt"

func ValidBinaryStringsWithCostLimit(n int, k int) []string {
	var ans []string
	var dfs func(pos int, prev byte, sum int, buf []byte)
	dfs = func(pos int, prev byte, sum int, buf []byte) {
		if sum > k {
			return
		}
		if pos == n {
			ans = append(ans, string(buf))
			return
		}

		// Place '0'
		buf[pos] = '0'
		dfs(pos+1, '0', sum, buf)

		// Place '1' only if no consecutive 1s
		if prev != '1' {
			buf[pos] = '1'
			dfs(pos+1, '1', sum+pos, buf)
		}
	}
	buf := make([]byte, n)
	dfs(0, '0', 0, buf)
	return ans
}

func main() {
	// Example 1
	fmt.Println(ValidBinaryStringsWithCostLimit(3, 1)) // Expected: ["000","010","100"]

	// Example 2
	fmt.Println(ValidBinaryStringsWithCostLimit(1, 0)) // Expected: ["0"]
}
```
