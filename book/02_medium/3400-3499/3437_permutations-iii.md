# 3437 — Permutations Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func permute(n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n!) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3437: Permutations III
// https://leetcode.com/problems/permutations-iii/
// Difficulty: Medium [Paid]
// Time: O(n!) Space: O(n)

import "fmt"

func permute(n int) [][]int {
	var ans [][]int
	used := make([]bool, n+1)
  // Alokasi slice integer
	cur := make([]int, 0, n)

	var dfs func()
	dfs = func() {
		if len(cur) == n {
  // Alokasi slice integer
			tmp := make([]int, n)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		start := len(cur)%2 + 1
		for i := start; i <= n; i += 2 {
			if !used[i] {
				used[i] = true
				cur = append(cur, i)
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return ans
}

func main() {
	fmt.Println(len(permute(3))) // 2
	fmt.Println(len(permute(4))) // 4
	for _, p := range permute(3) {
		fmt.Println(p)
	}
}
```
