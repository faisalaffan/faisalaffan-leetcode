# 3273 — Minimum Amount Of Damage Dealt To Bob

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minDamage(power int, damage []int, health []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3273: Minimum Amount of Damage Dealt to Bob
// https://leetcode.com/problems/minimum-amount-of-damage-dealt-to-bob/
// Difficulty: Hard
//
// Bob has `power` attack. Enemies have `damage[i]` and `health[i]`.
// Each second, Bob attacks one enemy (reducing its health by power),
// and every alive enemy deals its damage to Bob.
// Find the minimum total damage Bob takes by choosing the optimal kill order.
//
// This is a scheduling problem. The optimal order is to sort by
//   time_to_kill[i] / damage[i]
// where time_to_kill[i] = ceil(health[i] / power).
// Equivalently, compare using cross-multiplication:
//   t_i * d_j < t_j * d_i  => i before j

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minDamage(4, []int{1, 2, 3, 4}, []int{4, 5, 6, 8}))
	// Example 2
	fmt.Println(minDamage(1, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}))
	// Example 3
	fmt.Println(minDamage(10, []int{5, 5, 5}, []int{10, 20, 30}))
	// Example 4: single enemy
	fmt.Println(minDamage(3, []int{7}, []int{10}))
	// Example 5: large health
	fmt.Println(minDamage(2, []int{3, 4}, []int{10, 10}))
}

func minDamage(power int, damage []int, health []int) int64 {
	n := len(damage)
	type enemy struct {
		t int64 // time to kill (ceil(health/power))
		d int64 // damage per second
	}
	enemies := make([]enemy, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range enemies {
		enemies[i].t = int64((health[i] + power - 1) / power)
		enemies[i].d = int64(damage[i])
	}

	// Sort by t/d ratio ascending.
	// Equivalent sort comparator: t_i * d_j < t_j * d_i
  // Custom sort dengan comparator
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].t*enemies[j].d < enemies[j].t*enemies[i].d
	})

	var totalDamage int64
	var elapsed int64
	for _, e := range enemies {
		elapsed += e.t
		totalDamage += elapsed * e.d
	}

	return totalDamage
}
```
