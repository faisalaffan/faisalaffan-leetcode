# 2456 — Most Popular Video Creator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostPopularCreator(creators []string, ids []string, views []int) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2456: Most Popular Video Creator
// https://leetcode.com/problems/most-popular-video-creator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group by creator: track total views, best video (max views, smallest lexicographic).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	// [[bob, two], [alice, one]]
}

type Creator struct {
	total    int
	bestID   string
	bestView int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	creatorsMap := make(map[string]*Creator)
	maxTotal := 0

	for i, name := range creators {
		c, ok := creatorsMap[name]
		if !ok {
			c = &Creator{bestView: math.MinInt32}
			creatorsMap[name] = c
		}
		c.total += views[i]
		if views[i] > c.bestView || (views[i] == c.bestView && ids[i] < c.bestID) {
			c.bestView = views[i]
			c.bestID = ids[i]
		}
		if c.total > maxTotal {
			maxTotal = c.total
		}
	}

  // Membuat matriks/slice 2D untuk DP
	ans := make([][]string, 0)
	for name, c := range creatorsMap {
		if c.total == maxTotal {
			ans = append(ans, []string{name, c.bestID})
		}
	}
	return ans
}
```
