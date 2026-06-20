# 2105 — Watering Plants Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumRefill(plants []int, capacityA int, capacityB int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2105: Watering Plants II
// https://leetcode.com/problems/watering-plants-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	alice := 0
	bob := n - 1
	waterA := capacityA
	waterB := capacityB
	refills := 0

	for alice < bob {
		// Alice waters
		if waterA < plants[alice] {
			refills++
			waterA = capacityA
		}
		waterA -= plants[alice]
		alice++

		// Bob waters
		if waterB < plants[bob] {
			refills++
			waterB = capacityB
		}
		waterB -= plants[bob]
		bob--
	}

	// Same plant?
	if alice == bob {
		if waterA >= waterB {
			if waterA < plants[alice] {
				refills++
			}
		} else {
			if waterB < plants[bob] {
				refills++
			}
		}
	}

	return refills
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumRefill([]int{2, 2, 3, 3}, 5, 5))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimumRefill([]int{2, 2, 3, 3}, 3, 4))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minimumRefill([]int{5}, 10, 8))
	// Expected: 0
}
```
