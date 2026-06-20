# 0485 — Max Consecutive Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxConsecutiveOnes(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #485: Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}

func main() {
	fmt.Println(MaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(MaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
```
