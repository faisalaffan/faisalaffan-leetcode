# 3337 — Total Characters In String After Transformations Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthAfterTransformations(s string, t int, nums []int) int
```

> **💡 Hint:** Matrix exponentiation of the 26x26 transformation matrix. Since

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3337: Total Characters in String After Transformations II
// https://leetcode.com/problems/total-characters-in-string-after-transformations-ii/
// Difficulty: Hard
//
// Given a string s, each transformation replaces each character with a sequence
// of new characters defined by the nums array: character i (0-indexed from 'a')
// transforms into nums[i] + 1 characters (the next nums[i] letters cyclically).
// Count the total characters after t transformations (mod 1e9+7).
//
// Approach: Matrix exponentiation of the 26x26 transformation matrix. Since
// t can be large, use fast exponentiation (O(26^3 log t)).

import "fmt"

func main() {
	// Example 1
	fmt.Println(lengthAfterTransformations("ab", 2, []int{2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Example 2: single char, 1 transformation
	fmt.Println(lengthAfterTransformations("a", 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Example 3: no change
	fmt.Println(lengthAfterTransformations("a", 0, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Edge: all transforming
	fmt.Println(lengthAfterTransformations("z", 3, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}))
}

const MOD = 1000000007

func lengthAfterTransformations(s string, t int, nums []int) int {
	// Build initial count vector (size 26)
  // Alokasi slice integer
	cnt := make([]int64, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// Build transformation matrix M (26x26)
	// M[i][j] = 1 if character i transforms into character j
  // Membuat matriks/slice 2D untuk DP
	M := make([][]int64, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range M {
		M[i] = make([]int64, 26)
		for j := 1; j <= nums[i]; j++ {
			M[i][(i+j)%26] = 1
		}
	}

	// Compute M^t using fast exponentiation
	power := matPow(M, int64(t))

	// Result = cnt * M^t (as row vector)
  // Alokasi slice integer
	result := make([]int64, 26)
	for j := 0; j < 26; j++ {
		var sum int64
		for i := 0; i < 26; i++ {
			sum = (sum + cnt[i]*power[i][j]) % MOD
		}
		result[j] = sum
	}

	var ans int64
	for _, v := range result {
		ans = (ans + v) % MOD
	}
	return int(ans)
}

// matMul multiplies two 26x26 matrices
func matMul(a, b [][]int64) [][]int64 {
  // Membuat matriks/slice 2D untuk DP
	res := make([][]int64, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range res {
		res[i] = make([]int64, 26)
		for k := 0; k < 26; k++ {
			if a[i][k] == 0 {
				continue
			}
			for j := 0; j < 26; j++ {
				res[i][j] = (res[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}
	return res
}

// matPow computes matrix^exp
func matPow(mat [][]int64, exp int64) [][]int64 {
  // Membuat matriks/slice 2D untuk DP
	res := make([][]int64, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range res {
		res[i] = make([]int64, 26)
		res[i][i] = 1 // identity
	}

	base := mat
	for exp > 0 {
		if exp&1 == 1 {
			res = matMul(res, base)
		}
		base = matMul(base, base)
		exp >>= 1
	}
	return res
}
```
