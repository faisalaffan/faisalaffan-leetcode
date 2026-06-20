# 3209 — Number Of Subarrays With And Value Of K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64
```

> **💡 Hint:** maintain map of (AND value -> count) for subarrays ending at the

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3209: Number of Subarrays With AND Value of K
// https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/
// Difficulty: Hard
//
// Count subarrays whose bitwise AND equals exactly k.
// AND monotonically decreases as subarrays extend. For each ending position,
// there are at most O(log MAX) distinct AND values.
//
// Approach: maintain map of (AND value -> count) for subarrays ending at the
// current position.

import "fmt"

func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64 {
	var ans int64 = 0
  // Membuat map (HashMap) — pencarian O(1)
	cur := make(map[int]int)

	for _, x := range nums {
  // Membuat map (HashMap) — pencarian O(1)
		nxt := make(map[int]int)
		nxt[x] = 1
		for val, cnt := range cur {
			key := val & x
			nxt[key] += cnt
		}
		if cnt, ok := nxt[k]; ok {
			ans += int64(cnt)
		}
		cur = nxt
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubarraysWithAndValueOfK([]int{1, 1, 1}, 1)) // expect 6
	fmt.Println(numberOfSubarraysWithAndValueOfK([]int{1, 1, 2}, 1)) // expect ?
}
```
