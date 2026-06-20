# 2634 — Filter Elements From Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FilterElementsFromArray(arr []int, fn func(int, int) bool) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2634: Filter Elements from Array
// https://leetcode.com/problems/filter-elements-from-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Filters slice using a predicate.

import "fmt"

func main() {
	nums := []int{0, 10, 20, 30}
	greaterThan10 := func(n int, i int) bool { return n > 10 }
	fmt.Println(FilterElementsFromArray(nums, greaterThan10))

	nums2 := []int{1, 2, 3}
	firstIndex := func(n int, i int) bool { return i == 0 }
	fmt.Println(FilterElementsFromArray(nums2, firstIndex))
}

func FilterElementsFromArray(arr []int, fn func(int, int) bool) []int {
	result := []int{}
	for i, v := range arr {
		if fn(v, i) {
			result = append(result, v)
		}
	}
	return result
}
```
