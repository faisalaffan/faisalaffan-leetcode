# 1386 — Cinema Seat Allocation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxNumberOfFamilies(n int, reservedSeats [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Bitmask

**Kompleksitas Waktu:** O(n) where n = number of reserved seats  
**Kompleksitas Ruang:** O(k) where k = number of rows with reserved seats

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1386: Cinema Seat Allocation
// https://leetcode.com/problems/cinema-seat-allocation/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxNumberOfFamilies(3, [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}})) // 4

	// Test case 2
	fmt.Println(maxNumberOfFamilies(2, [][]int{{2, 1}, {1, 8}, {2, 6}})) // 2

	// Test case 3
	fmt.Println(maxNumberOfFamilies(4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}})) // 4
}

// Time: O(n) where n = number of reserved seats
// Space: O(k) where k = number of rows with reserved seats
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// Map row to bitmask of reserved seats (columns 2-9, bits 0-7)
  // Membuat map (HashMap) — pencarian O(1)
	rowMap := make(map[int]int)
	for _, seat := range reservedSeats {
		col := seat[1]
		if col >= 2 && col <= 9 {
			rowMap[seat[0]] |= 1 << (col - 2)
		}
	}

	total := (n - len(rowMap)) * 2 // rows with no reserved seats get 2 families each

	// Patterns for 4-person family groupings
	// Columns: 2-5 (left), 4-7 (middle), 6-9 (right)
	left := 0b11110000   // columns 2,3,4,5 (bits 0-3)
	middle := 0b00111100 // columns 4,5,6,7 (bits 2-5)
	right := 0b00001111  // columns 6,7,8,9 (bits 4-7)

	for _, mask := range rowMap {
		if mask&left == 0 && mask&right == 0 {
			total += 2
		} else if mask&left == 0 || mask&middle == 0 || mask&right == 0 {
			total += 1
		}
	}

	return total
}
```
