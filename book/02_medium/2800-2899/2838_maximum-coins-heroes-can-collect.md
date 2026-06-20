# 2838 — Maximum Coins Heroes Can Collect

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n log n + m log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2838: Maximum Coins Heroes Can Collect
// https://leetcode.com/problems/maximum-coins-heroes-can-collect/
// Difficulty: Medium [Paid]
// Time: O(n log n + m log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64 {
	n := len(monsters)
	type monster struct {
		power int
		coin  int
	}
	monsterList := make([]monster, n)
	for i := 0; i < n; i++ {
		monsterList[i] = monster{monsters[i], coins[i]}
	}
  // Custom sort dengan comparator
	sort.Slice(monsterList, func(i, j int) bool {
		return monsterList[i].power < monsterList[j].power
	})

  // Alokasi slice integer
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(monsterList[i].coin)
	}

  // Alokasi slice integer
	heroSorted := make([]int, len(heroes))
	copy(heroSorted, heroes)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(heroSorted)

  // Membuat map (HashMap) — pencarian O(1)
	heroMap := make(map[int]int64)
	for _, h := range heroSorted {
		if _, ok := heroMap[h]; ok {
			continue
		}
		idx := sort.Search(n, func(i int) bool {
			return monsterList[i].power > h
		})
		heroMap[h] = prefix[idx]
	}

  // Alokasi slice integer
	result := make([]int64, len(heroes))
	for i, h := range heroes {
		result[i] = heroMap[h]
	}
	return result
}

func main() {
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{1, 4, 2}, []int{1, 1, 3, 5}, []int{2, 3, 4, 5}))
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{5}, []int{1, 2, 3}, []int{10, 20, 30}))
}
```
