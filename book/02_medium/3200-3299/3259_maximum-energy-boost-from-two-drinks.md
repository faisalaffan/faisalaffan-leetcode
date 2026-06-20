# 3259 — Maximum Energy Boost From Two Drinks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3259: Maximum Energy Boost From Two Drinks
// https://leetcode.com/problems/maximum-energy-boost-from-two-drinks/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxEnergyBoost([]int{1, 3, 1}, []int{3, 1, 1}))       // 5
	fmt.Println(maxEnergyBoost([]int{4, 1, 1}, []int{1, 1, 3}))       // 7
	fmt.Println(maxEnergyBoost([]int{2, 2, 2, 2}, []int{3, 3, 3, 3})) // 12
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dpA, dpB := int64(0), int64(0)

	for i := 0; i < n; i++ {
		newA := max(dpA+int64(energyDrinkA[i]), dpB)
		newB := max(dpB+int64(energyDrinkB[i]), dpA)
		dpA, dpB = newA, newB
	}

	if dpA > dpB {
		return dpA
	}
	return dpB
}
```
