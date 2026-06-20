# 3334 — Find The Maximum Factor Score Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lcm(a, b int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3334: Find the Maximum Factor Score of Array
// https://leetcode.com/problems/find-the-maximum-factor-score-of-array/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxFactorScore([]int{2, 4, 8, 16})) // 64
	fmt.Println(maxFactorScore([]int{1, 2, 3, 4, 5})) // 60
	fmt.Println(maxFactorScore([]int{3}))             // 9
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func maxFactorScore(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

  // Alokasi slice integer
	preGCD := make([]int, n)
  // Alokasi slice integer
	preLCM := make([]int, n)
	preGCD[0] = nums[0]
	preLCM[0] = nums[0]
	for i := 1; i < n; i++ {
		preGCD[i] = gcd(preGCD[i-1], nums[i])
		preLCM[i] = lcm(preLCM[i-1], nums[i])
	}

  // Alokasi slice integer
	sufGCD := make([]int, n+1)
  // Alokasi slice integer
	sufLCM := make([]int, n+1)
	sufLCM[n] = 1
	for i := n - 1; i >= 0; i-- {
		sufGCD[i] = gcd(sufGCD[i+1], nums[i])
		sufLCM[i] = lcm(sufLCM[i+1], nums[i])
	}

	ans := int64(preGCD[n-1]) * int64(preLCM[n-1])

	for i := 0; i < n; i++ {
		var g int
		if i == 0 {
			g = sufGCD[1]
		} else {
			g = gcd(preGCD[i-1], sufGCD[i+1])
		}
		var l int
		if i == 0 {
			l = sufLCM[1]
		} else if i == n-1 {
			l = preLCM[n-2]
		} else {
			l = lcm(preLCM[i-1], sufLCM[i+1])
		}
		val := int64(g) * int64(l)
		if val > ans {
			ans = val
		}
	}

	return ans
}
```
