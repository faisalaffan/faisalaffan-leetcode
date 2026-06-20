# 2731 — Movement Of Robots

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MovementOfRobots(nums []int, s string, d int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2731: Movement of Robots
// https://leetcode.com/problems/movement-of-robots/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MovementOfRobots(nums []int, s string, d int) int {
	n := len(nums)
  // Alokasi slice integer
	pos := make([]int, n)
	for i, num := range nums {
		pos[i] = num
		if s[i] == 'R' {
			pos[i] += d
		} else {
			pos[i] -= d
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(pos)

	const mod = 1_000_000_007
	var sum int64
	var prefix int64
	for i, p := range pos {
		sum = (sum + int64(p)*int64(i) - prefix) % mod
		prefix += int64(p)
	}

	return int(sum)
}

func main() {
	fmt.Println(MovementOfRobots([]int{1, 0}, "RL", 2))
	fmt.Println(MovementOfRobots([]int{-2, 0, 2}, "RLL", 3))
}
```
