# 1747 — Leetflex Banned Accounts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findBanned(logins []Login) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Login struct {
	AccountID int
	IPAddress string
	LoginTime int
}

func findBanned(logins []Login) []int {
	// Group by account
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int][]Login)
	for _, l := range logins {
		groups[l.AccountID] = append(groups[l.AccountID], l)
	}

  // Membuat map (HashMap) — pencarian O(1)
	banned := make(map[int]bool)
	for accID, records := range groups {
  // Custom sort dengan comparator
		sort.Slice(records, func(i, j int) bool {
			return records[i].LoginTime < records[j].LoginTime
		})
		// Track latest login time per IP for this account
  // Membuat map (HashMap) — pencarian O(1)
		lastTime := make(map[string]int)
		for _, r := range records {
			if prevTime, ok := lastTime[r.IPAddress]; ok {
				// Same IP, update last time
				_ = prevTime
			}
			lastTime[r.IPAddress] = r.LoginTime
		}

		// Check if any IP has concurrent sessions
  // Membuat map (HashMap) — pencarian O(1)
		active := make(map[string]int)
		for _, r := range records {
			if _, ok := active[r.IPAddress]; ok {
				// Check if there's a different IP with an active session
				for ip, t := range active {
					if ip != r.IPAddress && t < r.LoginTime {
						banned[accID] = true
					}
				}
				// End current session for this IP and start new one
				delete(active, r.IPAddress)
			}
			active[r.IPAddress] = r.LoginTime
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, len(banned))
	for id := range banned {
		result = append(result, id)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	logins := []Login{
		{1, "IP1", 1},
		{1, "IP2", 2},
		{1, "IP1", 3},
	}
	fmt.Println(findBanned(logins))
}
```
