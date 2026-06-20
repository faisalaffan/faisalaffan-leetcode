# 1389 — Create Target Array In The Given Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func createTargetArray(nums []int, index []int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1389: Create Target Array in the Given Order
// https://leetcode.com/problems/create-target-array-in-the-given-order/
// Difficulty: Easy
//
// LeetCode submission: func createTargetArray(nums []int, index []int) []int

import "fmt"

func main() {
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})) // [0 4 1 3 2]
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0})) // [0 1 2 3 4]
}

// Time: O(n^2), Space: O(n)
func CreateTargetArrayInTheGivenOrder(nums []int, index []int) []int {
  // Alokasi slice integer
	res := make([]int, 0, len(nums))
	for i, idx := range index {
		res = append(res[:idx], append([]int{nums[i]}, res[idx:]...)...)
	}
	return res
}
```
