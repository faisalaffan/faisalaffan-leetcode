# 3731 — Find Missing Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMissingElements(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3731: Find Missing Elements
// https://leetcode.com/problems/find-missing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMissingElements([]int{1, 4, 2, 5}))
	fmt.Println(FindMissingElements([]int{7, 8, 6, 9}))
	fmt.Println(FindMissingElements([]int{5, 1}))
}

// Time: O(n)
// Space: O(n)
func FindMissingElements(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	has := make(map[int]bool)
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		has[v] = true
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}

  // Alokasi slice integer
	ans := make([]int, 0)
	for x := mn + 1; x < mx; x++ {
		if !has[x] {
			ans = append(ans, x)
		}
	}
	return ans
}
```
