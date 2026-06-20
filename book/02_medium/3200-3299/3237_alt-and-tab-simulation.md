# 3237 — Alt And Tab Simulation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func simulationResult(windows []int, queries []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3237: Alt and Tab Simulation
// https://leetcode.com/problems/alt-and-tab-simulation/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(n)

import (
	"container/list"
	"fmt"
)

func simulationResult(windows []int, queries []int) []int {
	order := list.New()
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]*list.Element)
	for _, w := range windows {
		e := order.PushBack(w)
		pos[w] = e
	}

	for _, q := range queries {
		if e, ok := pos[q]; ok {
			order.MoveToFront(e)
		}
	}

  // Alokasi slice integer
	ans := make([]int, 0, order.Len())
	for e := order.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(int))
	}
	return ans
}

func main() {
	fmt.Println(simulationResult([]int{1, 2, 3, 4}, []int{3, 1})) // Expected: [1 3 2 4]
	fmt.Println(simulationResult([]int{1, 2, 3}, []int{2}))        // Expected: [2 1 3]
}
```
