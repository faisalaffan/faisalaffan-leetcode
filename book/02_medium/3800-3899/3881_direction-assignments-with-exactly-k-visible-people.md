# 3881 — Direction Assignments With Exactly K Visible People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DirectionAssignmentsWithExactlyKVisiblePeople(n int, pos int, k int) int
```

> **💡 Hint:** Combinatorics. For each possible a (visible from left),

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3881: Direction Assignments with Exactly K Visible People
// https://leetcode.com/problems/direction-assignments-with-exactly-k-visible-people/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Combinatorics. For each possible a (visible from left),
// b = k - a must be visible from right. Sum 2 * C(pos, a) * C(n-pos-1, b).

import "fmt"

const MOD int64 = 1000000007

func modPow(a int64, b int, mod int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func DirectionAssignmentsWithExactlyKVisiblePeople(n int, pos int, k int) int {
	left := pos
	right := n - pos - 1

  // Alokasi slice integer
	fact := make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}

  // Alokasi slice integer
	invFact := make([]int64, n+1)
	invFact[n] = modPow(fact[n], int(MOD-2), MOD)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}

	nCr := func(nn, rr int) int64 {
		if rr < 0 || rr > nn {
			return 0
		}
		return fact[nn] * invFact[rr] % MOD * invFact[nn-rr] % MOD
	}

	var ans int64 = 0
	for a := 0; a <= left; a++ {
		b := k - a
		if b < 0 || b > right {
			continue
		}
		ans = (ans + nCr(left, a)*nCr(right, b)%MOD) % MOD
	}

	// Person at pos can be L or R
	ans = ans * 2 % MOD
	return int(ans)
}

func main() {
	// Example 1
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(3, 1, 0)) // Expected: 2

	// Example 2
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(3, 2, 1)) // Expected: 4

	// Example 3
	fmt.Println(DirectionAssignmentsWithExactlyKVisiblePeople(1, 0, 0)) // Expected: 2
}
```
