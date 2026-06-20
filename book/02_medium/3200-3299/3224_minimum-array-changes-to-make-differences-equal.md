# 3224 — Minimum Array Changes To Make Differences Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minChanges(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3224: Minimum Array Changes to Make Differences Equal
// https://leetcode.com/problems/minimum-array-changes-to-make-differences-equal/
// Difficulty: Medium
// Time: O(n + k) | Space: O(k)

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minChanges(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice integer
	diff := make([]int, k+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		curDiff := abs(a - b)

		// One change: can achieve any diff from 0 to max(a, b, k-a, k-b)
		maxReach := max(max(a, b), max(k-a, k-b))
		// One change can achieve any diff in [0, maxReach]
		diff[0]++
		if maxReach+1 <= k {
			diff[maxReach+1]--
		}

		// Zero changes: only curDiff
		diff[curDiff]--
		diff[curDiff+1]++
	}

	ans := n
	cur := 0
	for i := 0; i <= k; i++ {
		cur += diff[i]
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(minChanges([]int{1, 0, 1, 2, 4, 3}, 4)) // Expected: 2
	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6)) // Expected: ?
}
```
