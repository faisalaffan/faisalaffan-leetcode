# 3069 — Distribute Elements Into Two Arrays I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DistributeElementsIntoTwoArraysI(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3069: Distribute Elements Into Two Arrays I
// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: resultArray
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{2, 1, 3, 4})) // [2, 3, 4, 1]
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{5, 4, 3, 8})) // [5, 3, 4, 8]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: resultArray
func DistributeElementsIntoTwoArraysI(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}
```
