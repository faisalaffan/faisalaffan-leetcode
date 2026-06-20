# 0626 — Exchange Seats

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ExchangeSeats(students [][]interface{}) [][]interface
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #626: Exchange Seats
// https://leetcode.com/problems/exchange-seats/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Students: {id, name}
	students := [][]interface{}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "David"},
	}
	result := ExchangeSeats(students)
	for _, r := range result {
		fmt.Printf("id=%d, name=%s\n", r[0].(int), r[1].(string))
	}
}

func ExchangeSeats(students [][]interface{}) [][]interface{} {
	n := len(students)
  // Membuat matriks/slice 2D untuk DP
	result := make([][]interface{}, n)

	// Build map for easy lookup
  // Membuat map (HashMap) — pencarian O(1)
	studentMap := make(map[int]string)
	for _, student := range students {
		id := student[0].(int)
		name := student[1].(string)
		studentMap[id] = name
	}

	for id := 1; id <= n; id++ {
		if id%2 == 1 {
			if id+1 <= n {
				result[id-1] = []interface{}{id + 1, studentMap[id+1]}
			} else {
				result[id-1] = []interface{}{id, studentMap[id]}
			}
		} else {
			result[id-1] = []interface{}{id - 1, studentMap[id-1]}
		}
	}

	return result
}
```
