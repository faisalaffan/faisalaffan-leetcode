# 0350 — Intersection Of Two Arrays Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IntersectionOfTwoArraysIi(nums1, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n+m), Space: O(min(n,m))  
**Kompleksitas Ruang:** O(min(n,m))

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #350: Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(min(n,m))
func IntersectionOfTwoArraysIi(nums1, nums2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for _, v := range nums1 {
		count[v]++
	}
	var result []int
	for _, v := range nums2 {
		if count[v] > 0 {
			result = append(result, v)
			count[v]--
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArraysIi([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArraysIi([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```
