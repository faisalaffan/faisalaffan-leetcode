# 3767 — Maximize Points After Choosing K Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximizePointsAfterChoosingKTasks(technique1 []int, technique2 []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3767: Maximize Points After Choosing K Tasks
// https://leetcode.com/problems/maximize-points-after-choosing-k-tasks/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximizePointsAfterChoosingKTasks(technique1 []int, technique2 []int, k int) int64 {
	n := len(technique1)
	type task struct {
		diff   int
		t1, t2 int
	}
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		tasks[i] = task{
			diff: technique1[i] - technique2[i],
			t1:   technique1[i],
			t2:   technique2[i],
		}
	}

  // Custom sort
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].diff > tasks[j].diff
	})

	var ans int64
	posCount := 0
	for i := 0; i < n; i++ {
		if i < k {
			ans += int64(tasks[i].t1)
			posCount++
		} else if tasks[i].diff >= 0 {
			ans += int64(tasks[i].t1)
			posCount++
		} else {
			ans += int64(tasks[i].t2)
		}
	}
	return ans
}

func main() {
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{5, 3, 4}, []int{2, 6, 1}, 2))
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{1, 2, 3}, []int{4, 5, 6}, 1))
	fmt.Println(maximizePointsAfterChoosingKTasks([]int{10, 20}, []int{5, 15}, 1))
}
```
