# 3720 — Lexicographically Smallest Permutation Greater Than Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func lexicographicallySmallestPermutationGreaterThanTarget(s string, target string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n*26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3720: Lexicographically Smallest Permutation Greater Than Target
// https://leetcode.com/problems/lexicographically-smallest-permutation-greater-than-target/
// Difficulty: Medium
// Time: O(n*26) | Space: O(26)

import "fmt"

func lexicographicallySmallestPermutationGreaterThanTarget(s string, target string) string {
	n := len(s)
	var freq [26]int
	for i := 0; i < n; i++ {
		freq[s[i]-'a']++
	}

	var ans []byte

	var dfs func(idx int, check bool) bool
	dfs = func(idx int, check bool) bool {
		if idx == n {
			return check
		}
		for ch := 0; ch < 26; ch++ {
			if freq[ch] == 0 {
				continue
			}
			if !check && byte(ch)+'a' < target[idx] {
				continue
			}
			freq[ch]--
			ans = append(ans, byte(ch)+'a')
			nextCheck := check || byte(ch)+'a' > target[idx]
			if dfs(idx+1, nextCheck) {
				return true
			}
			ans = ans[:len(ans)-1]
			freq[ch]++
		}
		return false
	}

	if dfs(0, false) {
		return string(ans)
	}
	return ""
}

func main() {
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("abc", "bba"))
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("leet", "code"))
	fmt.Println(lexicographicallySmallestPermutationGreaterThanTarget("baba", "bbaa"))
}
```
