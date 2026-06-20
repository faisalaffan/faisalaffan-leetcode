# 3601 — Find Drivers With Improved Fuel Efficiency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindDriversWithImprovedFuelEfficiency(before, after []float64) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3601: Find Drivers with Improved Fuel Efficiency
// https://leetcode.com/problems/find-drivers-with-improved-fuel-efficiency/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	before := []float64{10, 15, 20}
	after := []float64{12, 14, 22}
	fmt.Println("Test 1:", FindDriversWithImprovedFuelEfficiency(before, after))
	// Test case 2
	before2 := []float64{10, 10}
	after2 := []float64{9, 11}
	fmt.Println("Test 2:", FindDriversWithImprovedFuelEfficiency(before2, after2))
	// Test case 3
	before3 := []float64{5}
	after3 := []float64{6}
	fmt.Println("Test 3:", FindDriversWithImprovedFuelEfficiency(before3, after3))
}

func FindDriversWithImprovedFuelEfficiency(before, after []float64) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(before) && i < len(after); i++ {
		if after[i] > before[i] {
			count++
		}
	}
	return count
}
```
