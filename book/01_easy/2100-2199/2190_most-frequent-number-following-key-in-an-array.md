# 2190 — Most Frequent Number Following Key In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2190: Most Frequent Number Following Key In an Array
// https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{1, 100, 200, 1, 100}, 1))    // 100
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{2, 2, 2, 2, 3}, 2))          // 2
}

// Time: O(n), Space: O(n)
func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}

	maxCount := 0
	result := 0
	for num, count := range freq {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	return result
}
```
