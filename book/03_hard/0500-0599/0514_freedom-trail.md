# 0514 — Freedom Trail

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findRotateSteps(ring string, key string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #514: Freedom Trail
// https://leetcode.com/problems/freedom-trail/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findRotateSteps("godding", "gd")) // Expected: 4
}

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	// pos[c] = list of indices in ring where character c appears
  // Membuat matriks/slice 2D untuk DP
	pos := make([][]int, 26)
	for i := 0; i < m; i++ {
		c := ring[i] - 'a'
		pos[c] = append(pos[c], i)
	}

	// dp[j] = min steps to spell up to current key char ending at ring index j
  // Alokasi slice integer
	dp := make([]int, m)
	for j := 0; j < m; j++ {
		if ring[j] == key[0] {
			dp[j] = minDist(j, 0, m) + 1 // steps to rotate + press
		} else {
			dp[j] = 1 << 30 // large number
		}
	}

	for i := 1; i < n; i++ {
  // Alokasi slice integer
		next := make([]int, m)
		for j := 0; j < m; j++ {
			next[j] = 1 << 30
		}
		for _, j := range pos[key[i]-'a'] {
			// from any previous position where we could have been
			for _, k := range pos[key[i-1]-'a'] {
				cost := dp[k] + minDist(j, k, m) + 1
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	ans := 1 << 30
	for j := 0; j < m; j++ {
		if dp[j] < ans {
			ans = dp[j]
		}
	}
	return ans
}

func minDist(i, j, m int) int {
	d := i - j
	if d < 0 {
		d = -d
	}
	if d > m-d {
		return m - d
	}
	return d
}
```
