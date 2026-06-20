# 0825 — Friends Of Appropriate Ages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FriendsOfAppropriateAges(ages []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + R) where R = 120  |  **Ruang:** O(R)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #825: Friends Of Appropriate Ages
// https://leetcode.com/problems/friends-of-appropriate-ages/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FriendsOfAppropriateAges([]int{16, 16}))
	fmt.Println(FriendsOfAppropriateAges([]int{16, 17, 18}))
	fmt.Println(FriendsOfAppropriateAges([]int{20, 30, 100, 110, 120}))
}

// Time: O(n + R) where R = 120 | Space: O(R)
func FriendsOfAppropriateAges(ages []int) int {
  // Alokasi slice
	cnt := make([]int, 121)
	for _, age := range ages {
		cnt[age]++
	}

  // Alokasi slice
	prefix := make([]int, 121)
	for i := 1; i <= 120; i++ {
		prefix[i] = prefix[i-1] + cnt[i]
	}

	ans := 0
	for age := 1; age <= 120; age++ {
		if cnt[age] == 0 {
			continue
		}
		left := age/2 + 8
		if left > age {
			continue
		}
		total := prefix[age] - prefix[left-1] - 1 // exclude self
		ans += cnt[age] * total
	}

	return ans
}
```
