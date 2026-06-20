# 1089 — Duplicate Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func duplicateZeros(arr []int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1089: Duplicate Zeros
// https://leetcode.com/problems/duplicate-zeros/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr1)
	fmt.Println(arr1) // [1,0,0,2,3,0,0,4]

	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2) // [1,2,3]
}

// LeetCode submission: duplicateZeros
func duplicateZeros(arr []int) {
	n := len(arr)
	possibleDups := 0
	for i := 0; i+possibleDups < n; i++ {
		if arr[i] == 0 {
			possibleDups++
		}
	}
	last := n - 1 - possibleDups
	for i := last; i >= 0; i-- {
		if i+possibleDups < n {
			arr[i+possibleDups] = arr[i]
		}
		if arr[i] == 0 {
			possibleDups--
			if i+possibleDups < n {
				arr[i+possibleDups] = 0
			}
		}
	}
}
```
