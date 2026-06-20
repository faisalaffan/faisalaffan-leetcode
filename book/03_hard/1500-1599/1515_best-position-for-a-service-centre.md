# 1515 — Best Position For A Service Centre

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getMinDistSum(positions [][]int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1515: Best Position for a Service Centre
// https://leetcode.com/problems/best-position-for-a-service-centre/
// Difficulty: Hard
//
// Weiszfeld's algorithm (iteratively reweighted) to find the geometric median.
// The geometric median minimizes the sum of Euclidean distances.
// Because the cost function is convex, gradient descent also works reliably.

func getMinDistSum(positions [][]int) float64 {
	n := len(positions)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	// Convert to float64 for precision
	pts := make([][2]float64, n)
	for i, p := range positions {
		pts[i] = [2]float64{float64(p[0]), float64(p[1])}
	}

	// Start at the centroid
	x, y := 0.0, 0.0
	for _, p := range pts {
		x += p[0]
		y += p[1]
	}
	x /= float64(n)
	y /= float64(n)

	// Weiszfeld iteration:
	//   x_{k+1} = (sum_i w_i * xi) / (sum_i w_i)
	//   where w_i = 1 / dist(p_i, current_point)
	// With safeguard for when the current point coincides with a data point.

	best := math.MaxFloat64

	for iter := 0; iter < 10000; iter++ {
		numerX, numerY, denom := 0.0, 0.0, 0.0
		for _, p := range pts {
			d := math.Sqrt((x-p[0])*(x-p[0]) + (y-p[1])*(y-p[1]))
			if d < 1e-12 {
				// We are (almost) exactly on a data point – gradient is undefined.
				// Perturb slightly and let the next iteration handle it.
				continue
			}
			w := 1.0 / d
			numerX += p[0] * w
			numerY += p[1] * w
			denom += w
		}
		if denom > 0 {
			x = numerX / denom
			y = numerY / denom
		}

		// Evaluate current sum of distances
		cur := 0.0
		for _, p := range pts {
			cur += math.Sqrt((x-p[0])*(x-p[0]) + (y-p[1])*(y-p[1]))
		}

		if cur < best {
			best = cur
		}

		// Stop when change is very small
		if iter > 0 && math.Abs(cur-best) < 1e-11 {
			break
		}
	}

	return best
}

func main() {
	// Example 1:
	// Input: positions = [[0,1],[1,0],[1,2],[2,1]]
	// Output: 4.00000
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}))

	// Example 2:
	// Input: positions = [[1,1],[3,3]]
	// Output: 2.82843
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{1, 1}, {3, 3}}))

	// Example 3:
	// Input: positions = [[1,1]]
	// Output: 0.00000
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{1, 1}}))
}
```
