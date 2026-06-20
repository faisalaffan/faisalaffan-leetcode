# 3284 — Sum Of Consecutive Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3284: Sum of Consecutive Subarrays
// https://leetcode.com/problems/sum-of-consecutive-subarrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(getSum([]int{1, 2, 3}))    // 20
	fmt.Println(getSum([]int{1, 2, 3, 5})) // 25
	fmt.Println(getSum([]int{1, 2, 3, 4})) // 50
}

func getSum(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	f, g := 1, 1
	s, t := nums[0], nums[0]
	ans := nums[0]

	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]

		if diff == 1 {
			f++
			s += f * nums[i]
			ans = (ans + s) % mod
		} else {
			f = 1
			s = nums[i]
		}

		if diff == -1 {
			g++
			t += g * nums[i]
			ans = (ans + t) % mod
		} else {
			g = 1
			t = nums[i]
		}

		if diff != 1 && diff != -1 {
			ans = (ans + nums[i]) % mod
		}
	}

	return ans
}
```
