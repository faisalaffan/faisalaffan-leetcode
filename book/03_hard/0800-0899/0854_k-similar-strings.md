# 0854 — K Similar Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func kSimilarity(s1 string, s2 string) int
```

> **💡 Hint:** BFS over string permutations. At each state, pick the first mismatching

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #854: K-Similar Strings
// https://leetcode.com/problems/k-similar-strings/
// Difficulty: Hard
// Approach: BFS over string permutations. At each state, pick the first mismatching
// position i and swap it with any later position j that would fix the mismatch.

import "fmt"

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	n := len(s1)
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[string]bool)
	queue := []string{s1}
	visited[s1] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			cur := queue[k]
			if cur == s2 {
				return steps
			}

			// Find first mismatching position
			i := 0
			for i < n && cur[i] == s2[i] {
				i++
			}

			// Try swapping with any position j > i where cur[j] == s2[i]
			// (this guarantees we fix position i)
			for j := i + 1; j < n; j++ {
				if cur[j] == s2[i] && cur[j] != s2[j] {
					next := []byte(cur)
					next[i], next[j] = next[j], next[i]
					str := string(next)
					if !visited[str] {
						visited[str] = true
						queue = append(queue, str)
					}
				}
			}
		}
		queue = queue[size:]
		steps++
	}

	return -1
}

func main() {
	fmt.Println(kSimilarity("ab", "ba"))       // Expected: 1
	fmt.Println(kSimilarity("abc", "bca"))      // Expected: 2
	fmt.Println(kSimilarity("abac", "baca"))    // Expected: 2
}
```
