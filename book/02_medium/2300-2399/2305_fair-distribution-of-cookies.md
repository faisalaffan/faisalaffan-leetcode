# 2305 — Fair Distribution Of Cookies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distributeCookies(cookies []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(k^n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2305: Fair Distribution of Cookies
// https://leetcode.com/problems/fair-distribution-of-cookies/
// Difficulty: Medium
// Time: O(k^n) | Space: O(k)

import "fmt"

func distributeCookies(cookies []int, k int) int {
  // Alokasi slice integer
	distribution := make([]int, k)
	minUnfairness := 1 << 30

	var dfs func(idx, maxSoFar int)
	dfs = func(idx, maxSoFar int) {
		if maxSoFar >= minUnfairness {
			return
		}
		if idx == len(cookies) {
			if maxSoFar < minUnfairness {
				minUnfairness = maxSoFar
			}
			return
		}
		for i := 0; i < k; i++ {
			distribution[i] += cookies[idx]
			newMax := maxSoFar
			if distribution[i] > newMax {
				newMax = distribution[i]
			}
			dfs(idx+1, newMax)
			distribution[i] -= cookies[idx]
			if distribution[i] == 0 {
				break
			}
		}
	}

	dfs(0, 0)
	return minUnfairness
}

func main() {
	// Test case 1
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	// Expected: 31

	// Test case 2
	fmt.Println(distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3))
	// Expected: 7
}
```
