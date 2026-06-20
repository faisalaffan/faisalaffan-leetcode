# 2956 — Find Common Elements Between Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindCommonElementsBetweenTwoArrays(nums1 []int, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2956: Find Common Elements Between Two Arrays
// https://leetcode.com/problems/find-common-elements-between-two-arrays/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findIntersectionValues
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6})) // [3, 4]
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{3, 4, 2, 3}, []int{1, 5})) // [0, 0]
}

// Time: O(n + m) | Space: O(n + m)
// LeetCode submission name: findIntersectionValues
func FindCommonElementsBetweenTwoArrays(nums1 []int, nums2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	set1 := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	set2 := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	count1 := 0
	for _, v := range nums1 {
		if set2[v] {
			count1++
		}
	}
	count2 := 0
	for _, v := range nums2 {
		if set1[v] {
			count2++
		}
	}
	return []int{count1, count2}
}
```
