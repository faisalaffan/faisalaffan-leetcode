# 3960 — Frequency Balance Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FrequencyBalanceSubarray(nums []int) int
```

> **💡 Hint:** Enumerate all subarrays. Track element frequencies and

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N^2)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3960: Frequency Balance Subarray
// https://leetcode.com/problems/frequency-balance-subarray/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Enumerate all subarrays. Track element frequencies and
// frequency-of-frequencies. Valid if: 1 distinct value, or exactly 2
// distinct freq values where one is double the other.

import "fmt"

func FrequencyBalanceSubarray(nums []int) int {
	n := len(nums)
	ans := 0

	for l := 0; l < n; l++ {
  // Membuat map (HashMap) — pencarian O(1)
		cnt := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
		freq := make(map[int]int) // frequency-of-frequencies

		for r := l; r < n; r++ {
			val := nums[r]
			oldF := cnt[val]
			if oldF > 0 {
				freq[oldF]--
				if freq[oldF] == 0 {
					delete(freq, oldF)
				}
			}
			newF := oldF + 1
			cnt[val] = newF
			freq[newF]++

			// Check validity
			valid := false
			if len(cnt) == 1 {
				valid = true
			} else if len(freq) == 2 {
				// Exactly two distinct frequency values
  // Alokasi slice integer
				vals := make([]int, 0, 2)
				for f := range freq {
					vals = append(vals, f)
				}
				a, b := vals[0], vals[1]
				if a > b {
					a, b = b, a
				}
				// One must be double the other
				if b == 2*a {
					valid = true
				}
			}

			if valid && (r-l+1) > ans {
				ans = r - l + 1
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 2, 1, 2, 3, 3, 3})) // Expected: 5

	// Example 2
	fmt.Println(FrequencyBalanceSubarray([]int{5, 5, 5, 5})) // Expected: 4

	// Example 3
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 3, 4})) // Expected: 1
}
```
