# 0347 — Top K Frequent Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func topKFrequent(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Count frequencies
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Bucket sort by frequency (index = frequency)
  // Membuat matriks/slice 2D untuk DP
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

  // Alokasi slice integer
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	// Expected: [1, 2]

	// Test case 2
	fmt.Println("Test 2:", topKFrequent([]int{1}, 1))
	// Expected: [1]

	// Test case 3
	fmt.Println("Test 3:", topKFrequent([]int{1, 2, 3, 1, 2, 1}, 2))
	// Expected: [1, 2]
}
```
