# 1349 — Maximum Students Taking Exam

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxStudents(seats [][]byte) int
```

> **💡 Hint:** Bitmask DP.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1349: Maximum Students Taking Exam
// https://leetcode.com/problems/maximum-students-taking-exam/
// Difficulty: Hard
//
// Approach: Bitmask DP.
// For each row generate all valid seating masks (no adjacent '1' bits,
// no overlap with broken seats). Then DP across rows: dp[row][mask] =
// max students in rows 0..row with mask on current row. Conflict check
// between rows: no diagonal adjacency (mask<<1 & prevMask == 0 and
// mask>>1 & prevMask == 0). Answer = max over masks on last row.

import "fmt"

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])

	// Row broken-seat masks (1 = broken, cannot sit there)
  // Alokasi slice integer
	broken := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seats[i][j] == '#' {
				broken[i] |= 1 << j
			}
		}
	}

	// Precompute all valid row masks (no adjacent 1s)
  // Alokasi slice integer
	valid := make([]int, 0, 1<<n)
	for mask := 0; mask < (1 << n); mask++ {
		if mask&(mask<<1) == 0 {
			valid = append(valid, mask)
		}
	}

	popcnt := func(x int) int {
		c := 0
		for x > 0 {
			c += x & 1
			x >>= 1
		}
		return c
	}

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 1<<n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	for _, mask := range valid {
		if mask&broken[0] == 0 {
			dp[0][mask] = popcnt(mask)
		}
	}

	for i := 1; i < m; i++ {
		for _, cur := range valid {
			if cur&broken[i] != 0 {
				continue
			}
			for _, prev := range valid {
				if prev&broken[i-1] != 0 {
					continue
				}
				if (cur<<1)&prev != 0 || (cur>>1)&prev != 0 {
					continue
				}
				if dp[i-1][prev] == -1 {
					continue
				}
				cnt := dp[i-1][prev] + popcnt(cur)
				if cnt > dp[i][cur] {
					dp[i][cur] = cnt
				}
			}
		}
	}

	ans := 0
	for _, mask := range valid {
		if dp[m-1][mask] > ans {
			ans = dp[m-1][mask]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxStudents([][]byte{
		{'#', '.', '#', '#', '.', '#'},
		{'.', '#', '#', '#', '#', '.'},
		{'#', '.', '#', '#', '.', '#'},
	})) // 4

	fmt.Println(maxStudents([][]byte{
		{'.', '#'},
		{'#', '#'},
		{'#', '.'},
	})) // 1

	fmt.Println(maxStudents([][]byte{
		{'.', '.'},
		{'.', '.'},
	})) // 4
}
```
