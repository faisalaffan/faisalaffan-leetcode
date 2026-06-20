# 2210 — Count Hills And Valleys In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountHillsAndValleysInAnArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2210: Count Hills and Valleys in an Array
// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountHillsAndValleysInAnArray([]int{2, 4, 1, 1, 6, 5})) // 3
	fmt.Println(CountHillsAndValleysInAnArray([]int{6, 6, 5, 5, 4, 1})) // 0
}

// Time: O(n), Space: O(1)
func CountHillsAndValleysInAnArray(nums []int) int {
	// Build flattened array (remove consecutive duplicates)
	flat := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			flat = append(flat, nums[i])
		}
	}

	count := 0
	for i := 1; i < len(flat)-1; i++ {
		if (flat[i] > flat[i-1] && flat[i] > flat[i+1]) || // hill
			(flat[i] < flat[i-1] && flat[i] < flat[i+1]) { // valley
			count++
		}
	}
	return count
}
```
