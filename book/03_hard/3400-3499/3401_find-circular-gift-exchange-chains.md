# 3401 — Find Circular Gift Exchange Chains

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindCircularGiftExchangeChains(n int, gifts [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3401: Find Circular Gift Exchange Chains
// https://leetcode.com/problems/find-circular-gift-exchange-chains/
// Difficulty: Hard [Paid]
//
// DFS cycle detection from each node.

import "fmt"

func main() {
	fmt.Println(FindCircularGiftExchangeChains(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}

func FindCircularGiftExchangeChains(n int, gifts [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, g := range gifts {
		u, v := g[0], g[1]
		adj[u] = append(adj[u], v)
	}

  // Alokasi slice integer
	visited := make([]int, n) // 0=unvisited, 1=in-stack, 2=done
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = 1
		count := 0
		for _, v := range adj[u] {
			if visited[v] == 1 {
				count++
			} else if visited[v] == 0 {
				count += dfs(v)
			}
		}
		visited[u] = 2
		return count
	}

	total := 0
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			total += dfs(i)
		}
	}
	return total
}
```
