# 0666 — Path Sum Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func pathSumIV(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #666: Path Sum IV
// https://leetcode.com/problems/path-sum-iv/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(pathSumIV([]int{113, 215, 221}))
	fmt.Println(pathSumIV([]int{113, 221}))
}

func pathSumIV(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	tree := make(map[int]int)
	for _, num := range nums {
		tree[num/10] = num % 10
	}

	total := 0
	var dfs func(key int, sum int)
	dfs = func(key int, sum int) {
		depth := key / 10
		pos := key % 10
		leftKey := (depth+1)*10 + pos*2 - 1
		rightKey := (depth+1)*10 + pos*2

		curSum := sum + tree[key]

		_, hasLeft := tree[leftKey]
		_, hasRight := tree[rightKey]

		if !hasLeft && !hasRight {
			total += curSum
			return
		}

		if hasLeft {
			dfs(leftKey, curSum)
		}
		if hasRight {
			dfs(rightKey, curSum)
		}
	}

	dfs(nums[0]/10, 0)
	return total
}
```
