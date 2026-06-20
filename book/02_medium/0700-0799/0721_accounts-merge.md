# 0721 — Accounts Merge

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func accountsMerge(accounts [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU), Merge Sort

**Kompleksitas Waktu:** O(nk * alpha(nk))  
**Kompleksitas Ruang:** O(nk)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #721: Accounts Merge
// https://leetcode.com/problems/accounts-merge/
// Difficulty: Medium
// Time: O(nk * alpha(nk))
// Space: O(nk)

import (
	"fmt"
	"sort"
)

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	parent := make(map[string]string)
  // Membuat map (HashMap) — pencarian O(1)
	owner := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]
		for _, email := range acc[1:] {
			parent[email] = email
			owner[email] = name
			union(firstEmail, email)
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		merged := append([]string{owner[root]}, emails...)
		result = append(result, merged)
	}

	return result
}
```
