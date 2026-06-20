# 2470 — Number Of Subarrays With Lcm Equal To K

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

**Kompleksitas Waktu:** O(n^2) worst-case  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2470: Number of Subarrays With LCM Equal to K
// https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k/
// Difficulty: Medium
// Time: O(n^2) worst-case | Space: O(1)
// For each start, expand and track LCM.

import "fmt"

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6)) // 4
	fmt.Println(subarrayLCM([]int{3}, 2))              // 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func subarrayLCM(nums []int, k int) int {
	ans := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		cur := 1
		for j := i; j < len(nums); j++ {
			cur = lcm(cur, nums[j])
			if cur == k {
				ans++
			}
			if cur > k {
				break
			}
		}
	}
	return ans
}
```
