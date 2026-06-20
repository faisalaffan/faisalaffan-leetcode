# 3717 — Minimum Operations To Make The Array Beautiful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToMakeTheArrayBeautiful(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3717: Minimum Operations to Make the Array Beautiful
// https://leetcode.com/problems/minimum-operations-to-make-the-array-beautiful/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToMakeTheArrayBeautiful(nums []int) int {
	ops := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]%prev != 0 {
			target := ((nums[i] / prev) + 1) * prev
			ops += target - nums[i]
			prev = target
		} else {
			prev = nums[i]
		}
	}
	return ops
}

func main() {
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{3, 7, 9}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{1, 1, 1}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{2, 3, 5}))
}
```
