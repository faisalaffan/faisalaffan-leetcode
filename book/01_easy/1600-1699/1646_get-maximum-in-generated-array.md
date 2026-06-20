# 1646 — Get Maximum In Generated Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetMaximumGenerated(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1646: Get Maximum in Generated Array
// https://leetcode.com/problems/get-maximum-in-generated-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GetMaximumGenerated(n int) int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
  // Alokasi slice integer
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	maxVal := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}

func main() {
	fmt.Println(GetMaximumGenerated(7))
	fmt.Println(GetMaximumGenerated(2))
	fmt.Println(GetMaximumGenerated(3))
}
```
