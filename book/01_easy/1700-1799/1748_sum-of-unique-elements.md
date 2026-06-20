# 1748 — Sum Of Unique Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfUnique(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1748: Sum of Unique Elements
// https://leetcode.com/problems/sum-of-unique-elements/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func SumOfUnique(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sum := 0
	for num, count := range freq {
		if count == 1 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(SumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(SumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(SumOfUnique([]int{1, 2, 3, 4, 5}))
}
```
