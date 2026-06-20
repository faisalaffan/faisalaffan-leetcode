# 0665 — Non Decreasing Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkPossibility(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #665: Non-decreasing Array
// https://leetcode.com/problems/non-decreasing-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	modified := false

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			}

			if i == 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}

			modified = true
		}
	}

	return true
}
```
