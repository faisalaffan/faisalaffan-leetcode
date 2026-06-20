# 3086 — Minimum Moves To Pick K Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumMoves(nums []int, k int) int64
```

> **💡 Hint:** Collect positions of 1s, use sliding window + prefix sum.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3086: Minimum Moves to Pick K Ones
// https://leetcode.com/problems/minimum-moves-to-pick-k-ones/
// Difficulty: Hard
// Time: O(n) | Space: O(n)
//
// Approach: Collect positions of 1s, use sliding window + prefix sum.
// For each window of size k, minimum moves = sum of distances to median.
// This is the classic "minimum moves to make array elements equal" = median minimizes L1 distance.

import (
	"fmt"
	"math"
)

func minimumMoves(nums []int, k int) int64 {
  // Alokasi slice integer
	pos := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			pos = append(pos, i)
		}
	}

	n := len(pos)
	if n < k {
		return -1
	}

  // Alokasi slice integer
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(pos[i])
	}

	var result int64 = math.MaxInt64

	for r := k; r <= n; r++ {
		l := r - k
		mid := l + k/2
		medianPos := int64(pos[mid])

		leftCount := int64(mid - l)
		leftSum := medianPos*leftCount - (prefix[mid] - prefix[l])

		rightCount := int64(r - mid - 1)
		rightSum := (prefix[r] - prefix[mid+1]) - medianPos*rightCount

		total := leftSum + rightSum
		if total < result {
			result = total
		}
	}

	return result
}

func main() {
	// Example 1
	fmt.Println("Test 1:", minimumMoves([]int{1, 0, 0, 1, 1, 0, 1}, 3))
	// Expected: 3

	// Example 2
	fmt.Println("Test 2:", minimumMoves([]int{1, 1, 0, 1}, 2))
	// Expected: 1

	// Example 3
	fmt.Println("Test 3:", minimumMoves([]int{1, 1, 1}, 2))
	// Expected: 1

	// All ones, pick all
	fmt.Println("Test 4:", minimumMoves([]int{1, 1, 1, 1, 1}, 5))
	// Expected: 6 (approx: median at index 2 (pos[2]=2). left: 2*2-(0+1)=4-1=3. right: (3+4)-2*2=7-4=3. total=6

	// Single one, k=1
	fmt.Println("Test 5:", minimumMoves([]int{0, 0, 1, 0, 0}, 1))
	// Expected: 0 (already at the position)

	// Not enough ones
	fmt.Println("Test 6:", minimumMoves([]int{0, 1, 0}, 5))
	// Expected: -1

	// Spaced out ones
	fmt.Println("Test 7:", minimumMoves([]int{1, 0, 0, 0, 1, 0, 0, 0, 1}, 2))
	// Expected: some value

	// Alternating pattern
	fmt.Println("Test 8:", minimumMoves([]int{1, 0, 1, 0, 1, 0, 1}, 3))
	// Pos: [0, 2, 4, 6], k=3.
	// Window 0-2: median=pos[1]=2, left=2-0=2*1-(0)=2, right=(4+6) - 3-1  hmm let it compute
}
```
