# 3476 — Maximize Profit From Task Assignment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeProfitFromTaskAssignment(difficulty []int, profit []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3476: Maximize Profit from Task Assignment
// https://leetcode.com/problems/maximize-profit-from-task-assignment/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximizeProfitFromTaskAssignment([]int{1, 3, 2, 4}, []int{2, 5, 3, 6}))
	// Test case 2
	fmt.Println("Test 2:", MaximizeProfitFromTaskAssignment([]int{5, 1, 3}, []int{10, 2, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximizeProfitFromTaskAssignment([]int{2}, []int{4}))
}

func MaximizeProfitFromTaskAssignment(difficulty []int, profit []int) int {
	type task struct {
		d int
		p int
	}
	n := len(difficulty)
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		tasks[i] = task{difficulty[i], profit[i]}
	}
  // Custom sort dengan comparator
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d || (tasks[i].d == tasks[j].d && tasks[i].p > tasks[j].p)
	})

	maxProfit := 0
	best := 0
	for i := 0; i < n; i++ {
		if tasks[i].p > best {
			best = tasks[i].p
		}
		maxProfit += best
	}
	return maxProfit
}
```
