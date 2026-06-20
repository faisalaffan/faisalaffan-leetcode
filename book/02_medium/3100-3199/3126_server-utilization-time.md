# 3126 — Server Utilization Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func serverUtilizationTime(logs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3126: Server Utilization Time
// https://leetcode.com/problems/server-utilization-time/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func serverUtilizationTime(logs [][]int) int {
	if len(logs) == 0 {
		return 0
	}

  // Custom sort dengan comparator
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	merged := [][]int{logs[0]}
	for i := 1; i < len(logs); i++ {
		last := merged[len(merged)-1]
		if logs[i][0] <= last[1] {
			if logs[i][1] > last[1] {
				last[1] = logs[i][1]
			}
		} else {
			merged = append(merged, logs[i])
		}
	}

	total := 0
	for _, seg := range merged {
		total += seg[1] - seg[0]
	}
	return total
}

func main() {
	fmt.Println(serverUtilizationTime([][]int{{0, 5}, {2, 7}, {8, 10}})) // Expected: 9
	fmt.Println(serverUtilizationTime([][]int{{1, 3}, {3, 5}}))          // Expected: 4
}
```
