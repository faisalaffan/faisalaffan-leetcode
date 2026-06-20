# 0134 — Gas Station

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canCompleteCircuit(gas []int, cost []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #134: Gas Station
// https://leetcode.com/problems/gas-station/
// Difficulty: Medium

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	start, tank := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}

func main() {
	// Test case 1
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3

	// Test case 2
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})) // -1

	// Test case 3
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) // 4
}

// Time: O(n) | Space: O(1)
```
