# 1934 — Confirmation Rate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConfirmationRate(signups [][]int, confirmations [][]int) []float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1934: Confirmation Rate
// https://leetcode.com/problems/confirmation-rate/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// signups: [user_id, signup_date]
	// confirmations: [user_id, action, confirmation_date]
	signups := [][]int{{1, 1}, {2, 1}, {3, 1}}
	confirmations := [][]int{{1, 1, 1}, {1, 1, 2}, {2, 1, 1}, {2, 0, 2}, {3, 0, 1}}
	fmt.Println(ConfirmationRate(signups, confirmations))
}

// Time: O(n + m), Space: O(n)
func ConfirmationRate(signups [][]int, confirmations [][]int) []float64 {
  // Membuat map (HashMap) — pencarian O(1)
	userMap := make(map[int][]int) // user_id -> [confirmed, total]
	for _, s := range signups {
		userMap[s[0]] = []int{0, 0}
	}
	for _, c := range confirmations {
		uid := c[0]
		if _, ok := userMap[uid]; ok {
			userMap[uid][1]++
			if c[1] == 1 { // confirmed
				userMap[uid][0]++
			}
		}
	}

	result := make([]float64, 0, len(signups))
	for _, s := range signups {
		uid := s[0]
		data := userMap[uid]
		if data[1] == 0 {
			result = append(result, 0.0)
		} else {
			result = append(result, float64(data[0])/float64(data[1]))
		}
	}
	return result
}
```
