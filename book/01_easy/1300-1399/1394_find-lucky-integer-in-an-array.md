# 1394 — Find Lucky Integer In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLucky(arr []int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1394: Find Lucky Integer in an Array
// https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func findLucky(arr []int) int

import "fmt"

func main() {
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 3, 4}))       // 2
	fmt.Println(FindLuckyIntegerInAnArray([]int{1, 2, 2, 3, 3, 3})) // 3
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 2, 3, 3}))    // -1
}

// Time: O(n), Space: O(n)
func FindLuckyIntegerInAnArray(arr []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int, len(arr))
	for _, v := range arr {
		freq[v]++
	}
	ans := -1
	for k, v := range freq {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
```
