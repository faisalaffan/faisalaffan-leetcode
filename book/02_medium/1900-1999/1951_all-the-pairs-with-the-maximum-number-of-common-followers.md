# 1951 — All The Pairs With The Maximum Number Of Common Followers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxCommonFollowers(relations [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2 * m) roughly, Space: O(n*m)  
**Kompleksitas Ruang:** O(n*m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1951: All the Pairs With the Maximum Number of Common Followers
// https://leetcode.com/problems/all-the-pairs-with-the-maximum-number-of-common-followers/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// relations: [user_id, follower_id]
	relations := [][]int{{1, 3}, {2, 3}, {3, 4}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(MaxCommonFollowers(relations))
}

// Time: O(n^2 * m) roughly, Space: O(n*m)
func MaxCommonFollowers(relations [][]int) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	followers := make(map[int]map[int]bool)
	for _, r := range relations {
		user, follower := r[0], r[1]
		if followers[user] == nil {
			followers[user] = make(map[int]bool)
		}
		followers[user][follower] = true
	}

  // Alokasi slice integer
	users := make([]int, 0, len(followers))
	for u := range followers {
		users = append(users, u)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(users)

	maxCommon := 0
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			common := 0
			for f := range followers[a] {
				if followers[b][f] {
					common++
				}
			}
			if common > maxCommon {
				maxCommon = common
				result = [][]int{{a, b}}
			} else if common == maxCommon && common > 0 {
				result = append(result, []int{a, b})
			}
		}
	}
	return result
}
```
