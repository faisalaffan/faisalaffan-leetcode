# 0349 — Intersection Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IntersectionOfTwoArrays(nums1, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n+m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #349: Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n)
func IntersectionOfTwoArrays(nums1, nums2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	var result []int
	for _, v := range nums2 {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```
