# 0697 — Degree Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findShortestSubArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #697: Degree of an Array
// https://leetcode.com/problems/degree-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))       // 2
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})) // 6
	fmt.Println(findShortestSubArray([]int{1}))                    // 1
}

// findShortestSubArray finds the smallest subarray length with the same degree as the array.
// Time: O(n). Space: O(n).
func findShortestSubArray(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	first := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	maxCount := 0
	minLen := len(nums)

	for i, v := range nums {
		if _, ok := first[v]; !ok {
			first[v] = i
		}
		count[v]++
		if count[v] > maxCount {
			maxCount = count[v]
		}
	}

	for v, c := range count {
		if c == maxCount {
			// find last occurrence
			last := 0
			for i := len(nums) - 1; i >= 0; i-- {
				if nums[i] == v {
					last = i
					break
				}
			}
			length := last - first[v] + 1
			if length < minLen {
				minLen = length
			}
		}
	}
	return minLen
}
```
