# 2229 — Check If An Array Is Consecutive

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfAnArrayIsConsecutive(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2229: Check if an Array Is Consecutive
// https://leetcode.com/problems/check-if-an-array-is-consecutive/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 4, 2})) // true
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 5}))    // false
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 4}))       // false
}

// Time: O(n), Space: O(n)
func CheckIfAnArrayIsConsecutive(nums []int) bool {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[int]bool)
	min, max := nums[0], nums[0]

	for _, v := range nums {
		set[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if max-min+1 != len(nums) {
		return false
	}

	for i := min; i <= max; i++ {
		if !set[i] {
			return false
		}
	}
	return true
}
```
