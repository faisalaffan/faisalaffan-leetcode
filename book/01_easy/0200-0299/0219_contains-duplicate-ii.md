# 0219 — Contains Duplicate Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ContainsNearbyDuplicate(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #219: Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsNearbyDuplicate(nums []int, k int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[n]; ok && i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}

func main() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
```
