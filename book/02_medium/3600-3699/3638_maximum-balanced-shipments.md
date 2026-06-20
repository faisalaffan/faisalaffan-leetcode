# 3638 — Maximum Balanced Shipments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxBalancedShipments(weight []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3638: Maximum Balanced Shipments
// https://leetcode.com/problems/maximum-balanced-shipments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxBalancedShipments(weight []int) int {
	ans := 0
	mx := 0
	for _, x := range weight {
		if x > mx {
			mx = x
		}
		if x < mx {
			ans++
			mx = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(maxBalancedShipments([]int{2, 5, 1, 4, 3}))
	fmt.Println(maxBalancedShipments([]int{4, 4}))
	fmt.Println(maxBalancedShipments([]int{1, 3, 2, 4, 5, 2}))
}
```
