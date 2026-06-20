# 0923 — 3Sum With Multiplicity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func threeSumMulti(arr []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1) if sort in-place considered O(1), else O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #923: 3Sum With Multiplicity
// https://leetcode.com/problems/3sum-with-multiplicity/
// Difficulty: Medium

import "fmt"

const mod = 1_000_000_007

// Time: O(n^2) | Space: O(1) if sort in-place considered O(1), else O(n)
func threeSumMulti(arr []int, target int) int {
	var cnt [101]int
	for _, v := range arr {
		cnt[v]++
	}

	ans := 0
	// Case 1: all three same
	for i := 0; i <= 100; i++ {
		if cnt[i] >= 3 && i*3 == target {
			ans = (ans + cnt[i]*(cnt[i]-1)*(cnt[i]-2)/6) % mod
		}
		// Case 2: two same, one different
		if cnt[i] >= 2 {
			remain := target - 2*i
			if remain >= 0 && remain <= 100 && remain != i && cnt[remain] > 0 {
				ans = (ans + cnt[i]*(cnt[i]-1)/2*cnt[remain]) % mod
			}
		}
		// Case 3: all three different
		for j := i + 1; j <= 100; j++ {
			k := target - i - j
			if k > j && k <= 100 && cnt[k] > 0 {
				ans = (ans + cnt[i]*cnt[j]*cnt[k]) % mod
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5))
	fmt.Println(threeSumMulti([]int{2, 1, 3}, 6))
}
```
