# 0740 — Delete And Earn

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func deleteAndEarn(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + k) where k is max value  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #740: Delete and Earn
// https://leetcode.com/problems/delete-and-earn/
// Difficulty: Medium
// Time: O(n + k) where k is max value
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}

  // Alokasi slice integer
	values := make([]int, maxVal+1)
	for _, n := range nums {
		values[n] += n
	}

	prev2, prev1 := 0, values[1]
	for i := 2; i <= maxVal; i++ {
		curr := max(prev1, prev2+values[i])
		prev2, prev1 = prev1, curr
	}

	return prev1
}
```
