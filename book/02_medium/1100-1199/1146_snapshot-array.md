# 1146 — Snapshot Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor(length int) SnapshotArray`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** set O(1) amortized, snap O(1), get O(log k) where k = history length  |  **Ruang:** O(n + total_sets)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1146: Snapshot Array
// https://leetcode.com/problems/snapshot-array/
// Difficulty: Medium

// SnapshotArray supports set(index, val), snap(), and get(index, snap_id).
// Instead of copying full array on each snap, store (snap_id, value) history per index.

// Time: set O(1) amortized, snap O(1), get O(log k) where k = history length
// Space: O(n + total_sets)

type SnapshotArray struct {
	data   [][]pair
	snapID int
}

type pair struct {
	snapID int
	val    int
}

func Constructor(length int) SnapshotArray {
  // Matriks 2D
	data := make([][]pair, length)
	for i := 0; i < length; i++ {
		data[i] = make([]pair, 0)
		// Initialize with snapID -1, value 0
		data[i] = append(data[i], pair{-1, 0})
	}
	return SnapshotArray{data: data, snapID: 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	idx := &this.data[index]
	// If last entry has same snapID, just update value
	if len(*idx) > 0 && (*idx)[len(*idx)-1].snapID == this.snapID {
		(*idx)[len(*idx)-1].val = val
		return
	}
	*idx = append(*idx, pair{this.snapID, val})
}

func (this *SnapshotArray) Snap() int {
	id := this.snapID
	this.snapID++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	history := this.data[index]
	// Binary search for the largest snapID <= snap_id
	i := sort.Search(len(history), func(i int) bool {
		return history[i].snapID > snap_id
	})
	return history[i-1].val
}

func main() {
	sa := Constructor(3)
	sa.Set(0, 5)
	fmt.Printf("snap: %d (expected: 0)\n", sa.Snap())
	sa.Set(0, 6)
	fmt.Printf("get(0, 0) = %d (expected: 5)\n", sa.Get(0, 0))
	fmt.Printf("get(0, 1) = %d (expected: 6)\n", sa.Get(0, 1))

	sa2 := Constructor(1)
	sa2.Set(0, 1)
	fmt.Printf("snap: %d (expected: 0)\n", sa2.Snap())
	fmt.Printf("get(0, 0) = %d (expected: 1)\n", sa2.Get(0, 0))
	sa2.Set(0, 2)
	fmt.Printf("snap: %d (expected: 1)\n", sa2.Snap())
	fmt.Printf("get(0, 0) = %d (expected: 1)\n", sa2.Get(0, 0))
	fmt.Printf("get(0, 1) = %d (expected: 2)\n", sa2.Get(0, 1))
}
```
