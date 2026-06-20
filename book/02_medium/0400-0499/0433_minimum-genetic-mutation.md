# 0433 — Minimum Genetic Mutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minMutation(startGene string, endGene string, bank []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** O(4^10) worst case (BFS)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #433: Minimum Genetic Mutation
// https://leetcode.com/problems/minimum-genetic-mutation/
// Difficulty: Medium
// Time: O(4^10) worst case (BFS) | Space: O(n)

import "fmt"

func minMutation(startGene string, endGene string, bank []string) int {
  // HashMap: O(1) lookup
	bankSet := make(map[string]bool)
	for _, b := range bank {
		bankSet[b] = true
	}

	if !bankSet[endGene] {
		return -1
	}

	genes := []byte{'A', 'C', 'G', 'T'}
	queue := []string{startGene}
	visited := map[string]bool{startGene: true}
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endGene {
				return steps
			}

			b := []byte(curr)
			for j := 0; j < 8; j++ {
				orig := b[j]
				for _, g := range genes {
					if g == orig {
						continue
					}
					b[j] = g
					next := string(b)
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				b[j] = orig
			}
		}
		steps++
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	// Expected: 2

	// Test case 3: Not in bank
	fmt.Println("Test 3:", minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
	// Expected: 3
}
```
