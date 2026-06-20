# 1460 — Make Two Arrays Equal By Reversing Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canBeEqual(target []int, arr []int) bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1460: Make Two Arrays Equal by Reversing Subarrays
// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func canBeEqual(target []int, arr []int) bool

import "fmt"

func main() {
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{1, 2, 3, 4}, []int{2, 4, 1, 3})) // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{7}, []int{7}))                    // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{3, 7, 9}, []int{3, 7, 11}))       // false
}

// Time: O(n), Space: O(n)
func MakeTwoArraysEqualByReversingSubarrays(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int, len(target))
	for _, v := range target {
		freq[v]++
	}
	for _, v := range arr {
		freq[v]--
		if freq[v] < 0 {
			return false
		}
	}
	return true
}
```
