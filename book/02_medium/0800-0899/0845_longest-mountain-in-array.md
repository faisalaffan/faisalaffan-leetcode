# 0845 — Longest Mountain In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestMountainInArray(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #845: Longest Mountain in Array
// https://leetcode.com/problems/longest-mountain-in-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LongestMountainInArray([]int{2, 1, 4, 7, 3, 2, 5}))
	fmt.Println(LongestMountainInArray([]int{2, 2, 2}))
	fmt.Println(LongestMountainInArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

// Time: O(n) | Space: O(1)
func LongestMountainInArray(arr []int) int {
	n := len(arr)
	ans := 0
	i := 1

	for i < n {
		// Skip non-increasing start
		for i < n && arr[i] == arr[i-1] {
			i++
		}

		// Climb up
		up := 0
		for i < n && arr[i] > arr[i-1] {
			up++
			i++
		}

		// Climb down
		down := 0
		for i < n && arr[i] < arr[i-1] {
			down++
			i++
		}

		// Valid mountain needs both up and down segments
		if up > 0 && down > 0 {
			if up+down+1 > ans {
				ans = up + down + 1
			}
		}
	}

	return ans
}
```
