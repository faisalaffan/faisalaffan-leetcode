# 2845 — Count Of Interesting Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountOfInterestingSubarrays(nums []int, modulo int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2845: Count of Interesting Subarrays
// https://leetcode.com/problems/count-of-interesting-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountOfInterestingSubarrays(nums []int, modulo int, k int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	prefix := make(map[int]int64)
	prefix[0] = 1
	var count int64
	var sum int

	for _, num := range nums {
		if num%modulo == k {
			sum++
		}
		need := (sum%modulo - k%modulo + modulo) % modulo
		count += prefix[need]
		prefix[sum%modulo]++
	}

	return count
}

func main() {
	fmt.Println(CountOfInterestingSubarrays([]int{3, 2, 4}, 2, 1))
	fmt.Println(CountOfInterestingSubarrays([]int{1, 2, 3, 4}, 3, 1))
}
```
