# 2127 — Maximum Employees To Be Invited To A Meeting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumInvitations(favorite []int) int
```

> **💡 Hint:** Cycle detection in functional graph.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Topological Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Topological Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2127: Maximum Employees to Be Invited to a Meeting
// https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/
// Difficulty: Hard
//
// Approach: Cycle detection in functional graph.
// Each node has exactly one outgoing edge (favorite[i]).
// Two cases contribute to the answer:
//   1. 2-cycles (mutual favorites): attach longest chain feeding into each node.
//   2. Cycles of length >= 3: must take the entire cycle, no chains attached.
//
// Use topological sort (Kahn's) to compute chain lengths for non-cycle nodes,
// then detect cycles among remaining nodes.

import "fmt"

func main() {
	// Example from problem statement
	favorite1 := []int{2, 2, 1, 2}
	fmt.Printf("maximumInvitations(%v) = %d (expected 3)\n", favorite1, maximumInvitations(favorite1))

	// Additional tests
	favorite2 := []int{1, 2, 0}
	fmt.Printf("maximumInvitations(%v) = %d (expected 3)\n", favorite2, maximumInvitations(favorite2))

	favorite3 := []int{1, 0, 3, 2}
	fmt.Printf("maximumInvitations(%v) = %d (expected 4)\n", favorite3, maximumInvitations(favorite3))

	favorite4 := []int{1, 0}
	fmt.Printf("maximumInvitations(%v) = %d (expected 2)\n", favorite4, maximumInvitations(favorite4))
}

func maximumInvitations(favorite []int) int {
	n := len(favorite)
  // Alokasi slice integer
	inDegree := make([]int, n)
	for _, f := range favorite {
		inDegree[f]++
	}

	// chainLen[i] = longest chain of non-cycle nodes ending at i
  // Alokasi slice integer
	chainLen := make([]int, n)
  // Alokasi slice integer
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	// Topological sort to compute chain lengths (removes non-cycle nodes)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		v := favorite[u]
		if chainLen[u]+1 > chainLen[v] {
			chainLen[v] = chainLen[u] + 1
		}
		inDegree[v]--
		if inDegree[v] == 0 {
			q = append(q, v)
		}
	}

	visited := make([]bool, n)
	totalChains := 0
	maxCycle := 0

	for i := 0; i < n; i++ {
		if inDegree[i] > 0 && !visited[i] {
			// Find the cycle
			cur := i
			cycleNodes := []int{}
			for !visited[cur] {
				visited[cur] = true
				cycleNodes = append(cycleNodes, cur)
				cur = favorite[cur]
			}

			cycleLen := len(cycleNodes)
			if cycleLen == 2 {
				a, b := cycleNodes[0], cycleNodes[1]
				totalChains += 2 + chainLen[a] + chainLen[b]
			} else {
				if cycleLen > maxCycle {
					maxCycle = cycleLen
				}
			}
		}
	}

	if totalChains > maxCycle {
		return totalChains
	}
	return maxCycle
}
```
