# 3207 — Maximum Points After Enemy Battles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumPoints(enemyEnergies []int, currentEnergy int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3207: Maximum Points After Enemy Battles
// https://leetcode.com/problems/maximum-points-after-enemy-battles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(enemyEnergies)
	n := len(enemyEnergies)
	energy := int64(currentEnergy)
	var points int64

	// First defeat all enemies we can defeat, accumulating energy
	i := 0
	for i < n && int64(enemyEnergies[i]) <= energy {
		energy -= int64(enemyEnergies[i])
		points++
		i++
	}

	if points == 0 {
		return 0
	}

	// Now we can use the smallest enemy to get energy
	smallest := enemyEnergies[0]
	for energy >= int64(smallest) {
		cnt := energy / int64(smallest)
		points += cnt
		energy %= int64(smallest)
		energy += int64(smallest) // Keep 1 point by defeating smallest again
		points--
	}
	return points
}

func main() {
	fmt.Println(maximumPoints([]int{3, 5, 6}, 3))  // Expected: 1
	fmt.Println(maximumPoints([]int{1, 2, 4, 8}, 3)) // Expected: 6
}
```
