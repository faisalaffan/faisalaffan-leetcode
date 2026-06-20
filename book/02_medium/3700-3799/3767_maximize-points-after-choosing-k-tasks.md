# 3767 — Maximize Points After Choosing K Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximizePointsAfterChoosingKTasks(technique1 []int, technique2 []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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

  // Custom sort dengan comparator
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
