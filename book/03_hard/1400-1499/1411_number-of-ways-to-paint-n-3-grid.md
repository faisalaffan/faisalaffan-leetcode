# 1411 — Number Of Ways To Paint N 3 Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func numOfWays(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1411: Number of Ways to Paint N × 3 Grid
// https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/
// Difficulty: Hard

import "fmt"

const mod1411 = 1_000_000_007

func numOfWays(n int) int {
	// Two pattern types for a 3-column row:
	// Pattern "ABA": 3 colors, first and third same (6 ways: 3*2)
	// Pattern "ABC": 3 colors, all different (6 ways: 3*2*1)
	aba, abc := 6, 6
	for i := 2; i <= n; i++ {
		// ABA can transition to:
		//   ABA: 3 ways (middle different from both ends)
		//   ABC: 2 ways (middle same as first, third different)
		// ABC can transition to:
		//   ABA: 2 ways (first and third same, middle different)
		//   ABC: 2 ways (all different, no color repeats position)
		newAba := (3*aba + 2*abc) % mod1411
		newAbc := (2*aba + 2*abc) % mod1411
		aba, abc = newAba, newAbc
	}
	return (aba + abc) % mod1411
}

func main() {
	// Example: n=1 -> 12
	fmt.Println(numOfWays(1))
}
```
