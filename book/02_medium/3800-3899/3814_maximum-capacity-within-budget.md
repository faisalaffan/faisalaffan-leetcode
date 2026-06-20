# 3814 — Maximum Capacity Within Budget

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumCapacityWithinBudget(costs []int, capacity []int, budget int) int
```

> **💡 Hint:** Sort by cost, use two pointers to find best pair with cost < budget.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Binary Search

**Kompleksitas Waktu:** O(N log N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3814: Maximum Capacity Within Budget
// https://leetcode.com/problems/maximum-capacity-within-budget/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N)
// Approach: Sort by cost, use two pointers to find best pair with cost < budget.

import (
	"fmt"
	"sort"
)

func MaximumCapacityWithinBudget(costs []int, capacity []int, budget int) int {
	n := len(costs)
	type machine struct {
		cost int
		cap  int
	}
	machines := make([]machine, n)
	for i := 0; i < n; i++ {
		machines[i] = machine{costs[i], capacity[i]}
	}

	// Sort by cost
  // Custom sort dengan comparator
	sort.Slice(machines, func(i, j int) bool {
		if machines[i].cost != machines[j].cost {
			return machines[i].cost < machines[j].cost
		}
		return machines[i].cap > machines[j].cap
	})

	ans := 0

	// Try single machine
	for _, m := range machines {
		if m.cost < budget && m.cap > ans {
			ans = m.cap
		}
	}

	// Try two machines using two pointers
	// For each machine, find best capacity at index < j with costs[i] + costs[j] < budget
	// Track max capacity seen so far for each cost
  // Membuat map (HashMap) — pencarian O(1)
	maxCapAtCost := make(map[int]int)
	for _, m := range machines {
		prevMax := 0
		if v, ok := maxCapAtCost[m.cost]; ok {
			prevMax = v
		}
		if m.cap > prevMax {
			maxCapAtCost[m.cost] = m.cap
		}
	}

  // Alokasi slice integer
	bestCap := make([]int, n)
	bestCap[0] = machines[0].cap
	for i := 1; i < n; i++ {
		if machines[i].cap > bestCap[i-1] {
			bestCap[i] = machines[i].cap
		} else {
			bestCap[i] = bestCap[i-1]
		}
	}

	for i := 1; i < n; i++ {
		// find a machine j < i with costs[j] + costs[i] < budget
		// binary search for largest cost < budget - costs[i]
		target := budget - machines[i].cost
		if target <= 0 {
			continue
		}
		lo, hi := 0, i-1
		best := -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if machines[mid].cost < target {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			total := machines[i].cap + bestCap[best]
			if total > ans {
				ans = total
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumCapacityWithinBudget([]int{4, 8, 5, 3}, []int{1, 5, 2, 7}, 8)) // Expected: 8

	// Example 2
	fmt.Println(MaximumCapacityWithinBudget([]int{3, 5, 7, 4}, []int{2, 4, 3, 6}, 7)) // Expected: 6

	// Example 3
	fmt.Println(MaximumCapacityWithinBudget([]int{2, 2, 2}, []int{3, 5, 4}, 5)) // Expected: 9
}
```
