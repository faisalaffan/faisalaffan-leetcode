# 0026 — Remove Duplicates From Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveDuplicates(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #26: Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveDuplicates(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{1, 1, 2}
	fmt.Println(RemoveDuplicates(n1), n1)
	n2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicates(n2), n2)
}
```
