# 0911 — Online Election

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(persons []int, times []int) TopVotedCandidate`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #911: Online Election
// https://leetcode.com/problems/online-election/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type TopVotedCandidate struct {
	times []int
	wins  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(times)
  // Alokasi slice
	wins := make([]int, n)
  // HashMap: O(1) lookup
	votes := make(map[int]int)
	leader := -1

	for i := 0; i < n; i++ {
		votes[persons[i]]++
		if leader == -1 || votes[persons[i]] >= votes[leader] {
			leader = persons[i]
		}
		wins[i] = leader
	}

	return TopVotedCandidate{times, wins}
}

func (this *TopVotedCandidate) Q(t int) int {
	idx := sort.SearchInts(this.times, t)
	if idx < len(this.times) && this.times[idx] == t {
		return this.wins[idx]
	}
	return this.wins[idx-1]
}

func main() {
	obj := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(obj.Q(3))
	fmt.Println(obj.Q(12))
	fmt.Println(obj.Q(25))
	fmt.Println(obj.Q(15))
	fmt.Println(obj.Q(24))
	fmt.Println(obj.Q(8))
}
```
