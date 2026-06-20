# 2804 — Array Prototype Foreach

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayPrototypeForeach(arr []int, callback func(int) )
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2804: Array Prototype ForEach
// https://leetcode.com/problems/array-prototype-foreach/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: JS problem, adapted to Go. Applies callback to each element.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	ArrayPrototypeForeach(nums, func(v int) {
		sum += v
	})
	fmt.Println(sum)
}

func ArrayPrototypeForeach(arr []int, callback func(int)) {
	for _, v := range arr {
		callback(v)
	}
}
```
