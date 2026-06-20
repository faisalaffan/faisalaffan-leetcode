# 2917 — Find The K Or Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheKOrOfAnArray(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n * 32)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2917: Find the K-or of an Array
// https://leetcode.com/problems/find-the-k-or-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findKOr
	fmt.Println(FindTheKOrOfAnArray([]int{7, 12, 9, 8, 9, 15}, 4)) // 9
	fmt.Println(FindTheKOrOfAnArray([]int{2, 12, 1, 11, 4, 5}, 6)) // 0
	fmt.Println(FindTheKOrOfAnArray([]int{10, 8, 5, 9, 11, 6, 8}, 1)) // 15
}

// Time: O(n * 32) | Space: O(1)
// LeetCode submission name: findKOr
func FindTheKOrOfAnArray(nums []int, k int) int {
	result := 0
	for bit := 0; bit < 32; bit++ {
		count := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				count++
			}
		}
		if count >= k {
			result |= (1 << bit)
		}
	}
	return result
}
```
