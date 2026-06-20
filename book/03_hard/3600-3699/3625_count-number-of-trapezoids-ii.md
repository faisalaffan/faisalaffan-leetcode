# 3625 — Count Number Of Trapezoids Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countTrapezoids(points [][]int) int
```

> **💡 Hint:** Group pairs by slope, count parallel combinations, subtract

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3625: Count Number of Trapezoids II
// https://leetcode.com/problems/count-number-of-trapezoids-ii/
// Difficulty: Hard
//
// Given points on a plane, count the number of unique trapezoids (convex
// quadrilaterals with at least one pair of parallel sides) that can be formed.
//
// Approach: Group pairs by slope, count parallel combinations, subtract
// parallelograms (which have both pairs parallel, counted twice).

import "fmt"

func main() {
	// Example 1
	fmt.Println(countTrapezoids([][]int{{0, 0}, {1, 1}, {2, 0}, {3, 1}}))
	// Example 2
	fmt.Println(countTrapezoids([][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}))
	// Edge: minimum points
	fmt.Println(countTrapezoids([][]int{{0, 0}, {1, 0}, {0, 1}, {2, 2}}))
}

func countTrapezoids(points [][]int) int {
	n := len(points)
	if n < 4 {
		return 0
	}

	// Group pairs by slope
	type pair struct{ dx, dy int }
	// Normalized slope representation (dx, dy) where gcd(dx,dy)=1, dx>0 or dx=0,dy>0
	normalize := func(dx, dy int) pair {
		if dx < 0 {
			dx, dy = -dx, -dy
		} else if dx == 0 {
			dy = 1
		}
		g := gcd(abs(dx), abs(dy))
		dx /= g
		dy /= g
		return pair{dx, dy}
	}

	// Count parallel pairs per slope
	type pairInfo struct {
		midX, midY int // midpoint for parallelogram detection
	}
  // Membuat map (HashMap) — pencarian O(1)
	slopePairs := make(map[pair][]pairInfo)
	parallelCount := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			s := normalize(dx, dy)
			mx := points[i][0] + points[j][0]
			my := points[i][1] + points[j][1]
			slopePairs[s] = append(slopePairs[s], pairInfo{mx, my})
		}
	}

	// Count trapezoids: for each slope, choose 2 pairs (C(cnt, 2))
	for _, pairs := range slopePairs {
		cnt := len(pairs)
		if cnt >= 2 {
			parallelCount += cnt * (cnt - 1) / 2
		}
	}

	// Count parallelograms (both pairs parallel) - they are counted twice
	// For each slope pair (i,j) and (p,q) with same midpoint, they form a parallelogram
	parallelogramCount := 0
	for _, pairs := range slopePairs {
  // Membuat map (HashMap) — pencarian O(1)
		midCount := make(map[pair]int)
		for _, p := range pairs {
			mp := pair{p.midX, p.midY}
			midCount[mp]++
		}
		for _, cnt := range midCount {
			if cnt >= 2 {
				parallelogramCount += cnt * (cnt - 1) / 2
			}
		}
	}

	// Trapezoids = parallel pairs - 2 * parallelograms
	// (each parallelogram has 2 pairs of parallel sides, counted twice in parallelCount)
	result := parallelCount - 2*parallelogramCount
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
