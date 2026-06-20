# 3880 — Minimum Absolute Difference Between Two Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAbsoluteDifferenceBetweenTwoValues(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3880: Minimum Absolute Difference Between Two Values
// https://leetcode.com/problems/minimum-absolute-difference-between-two-values/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumAbsoluteDifferenceBetweenTwoValues([]int{1, 0, 0, 2, 0, 1}))
	fmt.Println(MinimumAbsoluteDifferenceBetweenTwoValues([]int{1, 0, 1, 0}))
}

// Time: O(n)
// Space: O(1)
func MinimumAbsoluteDifferenceBetweenTwoValues(nums []int) int {
	last1, last2 := -1, -1
	ans := -1
	for i, v := range nums {
		if v == 1 {
			last1 = i
			if last2 != -1 {
				diff := i - last2
				if ans == -1 || diff < ans {
					ans = diff
				}
			}
		} else if v == 2 {
			last2 = i
			if last1 != -1 {
				diff := i - last1
				if ans == -1 || diff < ans {
					ans = diff
				}
			}
		}
	}
	return ans
}
```
