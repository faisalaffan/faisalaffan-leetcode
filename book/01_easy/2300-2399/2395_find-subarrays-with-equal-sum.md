# 2395 — Find Subarrays With Equal Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSubarraysWithEqualSum(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2395: Find Subarrays With Equal Sum
// https://leetcode.com/problems/find-subarrays-with-equal-sum/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(FindSubarraysWithEqualSum([]int{4, 2, 4}))   // true
	fmt.Println(FindSubarraysWithEqualSum([]int{1, 2, 3, 4, 5})) // false
}

func FindSubarraysWithEqualSum(nums []int) bool {
	seen := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
```
