# 3044 — Most Frequent Prime

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostFrequentPrime(mat [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(m*n*maxLen)  
**Kompleksitas Ruang:** O(K)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3044: Most Frequent Prime
// https://leetcode.com/problems/most-frequent-prime/
// Difficulty: Medium
// Time: O(m*n*maxLen) | Space: O(K)

import "fmt"

func main() {
	fmt.Println(mostFrequentPrime([][]int{{1, 1}, {9, 9}, {1, 1}}))
	fmt.Println(mostFrequentPrime([][]int{{7}}))
}

func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	cnt := map[int]int{}
	isPrime := func(x int) bool {
		if x < 2 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, d := range dirs {
				val := 0
				x, y := i, j
				for x >= 0 && x < m && y >= 0 && y < n {
					val = val*10 + mat[x][y]
					if val > 10 && isPrime(val) {
						cnt[val]++
					}
					x += d[0]
					y += d[1]
				}
			}
		}
	}
	maxCnt, maxVal := 0, -1
	for v, c := range cnt {
		if c > maxCnt || (c == maxCnt && v > maxVal) {
			maxCnt = c
			maxVal = v
		}
	}
	return maxVal
}
```
