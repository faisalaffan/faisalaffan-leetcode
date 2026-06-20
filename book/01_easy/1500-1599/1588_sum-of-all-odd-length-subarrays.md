# 1588 — Sum Of All Odd Length Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOddLengthSubarrays(arr []int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1588: Sum of All Odd Length Subarrays
// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func sumOddLengthSubarrays(arr []int) int

import "fmt"

func main() {
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 4, 2, 5, 3})) // 58
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 2}))          // 3
	fmt.Println(SumOfAllOddLengthSubarrays([]int{10, 11, 12}))    // 66
}

// Time: O(n), Space: O(1)
func SumOfAllOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0
	for i, v := range arr {
		contribution := ((i+1)*(n-i) + 1) / 2
		sum += v * contribution
	}
	return sum
}
```
