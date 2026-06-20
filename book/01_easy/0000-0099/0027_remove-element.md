# 0027 — Remove Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveElement(nums []int, val int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #27: Remove Element
// https://leetcode.com/problems/remove-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveElement(nums []int, val int) int {
	k := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{3, 2, 2, 3}
	fmt.Println(RemoveElement(n1, 3), n1)
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(RemoveElement(n2, 2), n2)
}
```
