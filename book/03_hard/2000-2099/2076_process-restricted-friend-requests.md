# 2076 — Process Restricted Friend Requests

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func newUnionFind(n int) *unionFind
```

> **💡 Hint:** Union-Find + Restriction Check

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2076: Process Restricted Friend Requests
// https://leetcode.com/problems/process-restricted-friend-requests/
// Difficulty: Hard
// Approach: Union-Find + Restriction Check

import "fmt"

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
  // Alokasi slice integer
	parent := make([]int, n)
  // Alokasi slice integer
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := newUnionFind(n)
	ans := make([]bool, len(requests))

	for i, req := range requests {
		u, v := req[0], req[1]
		ru, rv := uf.find(u), uf.find(v)
		canFriend := true

		// Check all restrictions: if both u's group and v's group
		// would contain both ends of a restriction, reject
		for _, res := range restrictions {
			a, b := res[0], res[1]
			ra, rb := uf.find(a), uf.find(b)
			// After union, ru and rv would be merged.
			// If (ra, rb) == (ru, rv) or (ra, rb) == (rv, ru), restriction is violated
			if (ra == ru && rb == rv) || (ra == rv && rb == ru) {
				canFriend = false
				break
			}
		}

		if canFriend {
			uf.union(u, v)
		}
		ans[i] = canFriend
	}

	return ans
}

func main() {
	fmt.Println("2076. Process Restricted Friend Requests")

	// Example 1
	n1 := 3
	restrictions1 := [][]int{{0, 1}}
	requests1 := [][]int{{0, 2}, {2, 1}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false])\n",
		n1, restrictions1, requests1, friendRequests(n1, restrictions1, requests1))

	// Example 2
	n2 := 5
	restrictions2 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	requests2 := [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false, true, false])\n",
		n2, restrictions2, requests2, friendRequests(n2, restrictions2, requests2))
}
```
