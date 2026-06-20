# 0932 — Beautiful Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulArray(n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS, Dynamic Programming

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #932: Beautiful Array
// https://leetcode.com/problems/beautiful-array/
// Difficulty: Medium

import "fmt"

// Time: O(n log n) | Space: O(n)
func beautifulArray(n int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[int][]int)
	var dfs func(int) []int
	dfs = func(n int) []int {
		if v, ok := memo[n]; ok {
			return v
		}
  // Alokasi slice integer
		res := make([]int, n)
		if n == 1 {
			res[0] = 1
		} else {
			left := dfs((n + 1) / 2)
			right := dfs(n / 2)
			for i, v := range left {
				res[i] = 2*v - 1
			}
			for i, v := range right {
				res[(n+1)/2+i] = 2 * v
			}
		}
		memo[n] = res
		return res
	}
	return dfs(n)
}

func main() {
	fmt.Println(beautifulArray(4))
	fmt.Println(beautifulArray(5))
	fmt.Println(beautifulArray(1))
}
```
