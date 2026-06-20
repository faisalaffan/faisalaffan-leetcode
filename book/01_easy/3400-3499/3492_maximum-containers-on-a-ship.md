# 3492 — Maximum Containers On A Ship

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumContainersOnAShip(n int, w int, maxWeight int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3492: Maximum Containers on a Ship
// https://leetcode.com/problems/maximum-containers-on-a-ship/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumContainersOnAShip(3, 2, 10))
	fmt.Println(MaximumContainersOnAShip(5, 3, 50))
}

// MaximumContainersOnAShip returns the maximum number of containers that can be loaded on an n x n ship, each weighing w, within maxWeight.
// Time: O(1). Space: O(1).
func MaximumContainersOnAShip(n int, w int, maxWeight int) int {
	maxBySpace := n * n
	maxByWeight := maxWeight / w
	if maxBySpace < maxByWeight {
		return maxBySpace
	}
	return maxByWeight
}
```
