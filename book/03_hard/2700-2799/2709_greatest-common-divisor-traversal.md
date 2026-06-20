# 2709 — Greatest Common Divisor Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canTraverseAllPairs(nums []int) bool
```

> **💡 Hint:** Union-Find with prime factorization via smallest prime factor (SPF) sieve.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2709: Greatest Common Divisor Traversal
// https://leetcode.com/problems/greatest-common-divisor-traversal/
// Difficulty: Hard
//
// Approach: Union-Find with prime factorization via smallest prime factor (SPF) sieve.
// For each number, factorize using SPF and union the number's index with each
// unique prime factor. After processing all numbers, check that all indices
// belong to the same connected component.

import "fmt"

func main() {
	// Example 1: [2,3,6] -> true
	fmt.Println(canTraverseAllPairs([]int{2, 3, 6}))
	// Example 2: [3,9,5] -> false
	fmt.Println(canTraverseAllPairs([]int{3, 9, 5}))
	// Example 3: [4,3,12,8] -> true
	fmt.Println(canTraverseAllPairs([]int{4, 3, 12, 8}))
}

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	// Find max value for sieve size
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	if maxVal < 2 {
		return false // all values are 1, can't connect
	}

	// Smallest Prime Factor sieve
  // Alokasi slice integer
	spf := make([]int, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		if spf[i] == 0 {
			for j := i; j <= maxVal; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	// Union-Find: indices 0..n-1 for array elements, n..n+maxPrimes for prime nodes
	// Instead of mapping primes, we use offset: prime p maps to n+p
	total := n + maxVal + 1
  // Alokasi slice integer
	parent := make([]int, total)
  // Alokasi slice integer
	size := make([]int, total)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	for i, v := range nums {
		if v == 1 {
			return false // 1 shares no prime factors with any other number
		}
		x := v
		prevPrime := 0
		for x > 1 {
			p := spf[x]
			if p != prevPrime {
				union(i, n+p)
				prevPrime = p
			}
			x /= p
		}
	}

	root := find(0)
	for i := 1; i < n; i++ {
		if find(i) != root {
			return false
		}
	}
	return true
}
```
