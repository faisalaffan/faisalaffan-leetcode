# 2439 — Minimize Maximum Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimizeArrayValue(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2439: Minimize Maximum of Array
// https://leetcode.com/problems/minimize-maximum-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Prefix average approach: we can distribute value to the left.
// The min possible max is the max prefix average (ceil).

import "fmt"

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6})) // 5
	fmt.Println(minimizeArrayValue([]int{10, 1}))      // 10
}

func minimizeArrayValue(nums []int) int {
	var sum int64
	ans := 0
	for i, v := range nums {
		sum += int64(v)
		avg := int((sum + int64(i)) / int64(i+1)) // ceil division
		if avg > ans {
			ans = avg
		}
	}
	return ans
}
```
