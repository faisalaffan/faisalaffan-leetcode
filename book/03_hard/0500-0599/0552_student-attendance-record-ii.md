# 0552 — Student Attendance Record Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkRecord(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #552: Student Attendance Record II
// https://leetcode.com/problems/student-attendance-record-ii/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func main() {
	fmt.Println(checkRecord(2))     // Expected: 8
	fmt.Println(checkRecord(10101)) // Expected: 183236316
}

func checkRecord(n int) int {
	// States:
	// dp[i][a][l] where:
	//   i = length
	//   a = number of absences (0 or 1)
	//   l = consecutive lates (0, 1, or 2)
	//
	// We use 3D DP. dp[a][l] = count for current length.

	prev := [2][3]int{}
	prev[0][0] = 1 // "" -> empty string

	for i := 0; i < n; i++ {
		cur := [2][3]int{}
		for a := 0; a <= 1; a++ {
			for l := 0; l <= 2; l++ {
				if prev[a][l] == 0 {
					continue
				}
				// Append 'P' (present) — resets lates
				cur[a][0] = (cur[a][0] + prev[a][l]) % mod
				// Append 'A' (absent)
				if a < 1 {
					cur[a+1][0] = (cur[a+1][0] + prev[a][l]) % mod
				}
				// Append 'L' (late)
				if l < 2 {
					cur[a][l+1] = (cur[a][l+1] + prev[a][l]) % mod
				}
			}
		}
		prev = cur
	}

	ans := 0
	for a := 0; a <= 1; a++ {
		for l := 0; l <= 2; l++ {
			ans = (ans + prev[a][l]) % mod
		}
	}
	return ans
}
```
