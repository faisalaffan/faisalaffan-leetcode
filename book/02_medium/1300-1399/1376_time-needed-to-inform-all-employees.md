# 1376 — Time Needed To Inform All Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numOfMinutes(n int, headID int, manager []int, informTime []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** O(n) where n = number of employees  
**Kompleksitas Ruang:** O(n) for memoization and adjacency list

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1376: Time Needed to Inform All Employees
// https://leetcode.com/problems/time-needed-to-inform-all-employees/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0})) // 1

	// Test case 2
	fmt.Println(numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1})) // 21

	// Test case 3
	fmt.Println(numOfMinutes(1, 0, []int{-1}, []int{0})) // 0
}

// Time: O(n) where n = number of employees
// Space: O(n) for memoization and adjacency list
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	// Build adjacency list (subordinates)
  // Membuat matriks/slice 2D untuk DP
	subordinates := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] != -1 {
			subordinates[manager[i]] = append(subordinates[manager[i]], i)
		}
	}

	// DFS with memoization
  // Alokasi slice integer
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(id int) int {
		if memo[id] != 0 {
			return memo[id]
		}
		maxTime := 0
		for _, sub := range subordinates[id] {
			time := dfs(sub)
			if time > maxTime {
				maxTime = time
			}
		}
		memo[id] = informTime[id] + maxTime
		return memo[id]
	}

	return dfs(headID)
}
```
