# 3384 — Team Dominance By Pass Success

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TeamDominanceByPassSuccess(passes [][]int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3384: Team Dominance by Pass Success
// https://leetcode.com/problems/team-dominance-by-pass-success/
// Difficulty: Hard [Paid]
//
// Aggregate pass stats per team: total passes, successful passes, dominance.

import "fmt"

func main() {
	fmt.Println(TeamDominanceByPassSuccess([][]int{{1, 1, 0}, {1, 1, 1}, {2, 0, 0}}))
}

func TeamDominanceByPassSuccess(passes [][]int) float64 {
	type teamStat struct{ total, succ int }
  // Membuat map (HashMap) — pencarian O(1)
	teams := make(map[int]*teamStat)

	for _, p := range passes {
		team, succ := p[0], p[1]
		if _, ok := teams[team]; !ok {
			teams[team] = &teamStat{}
		}
		teams[team].total++
		if succ == 1 {
			teams[team].succ++
		}
	}

	best := 0.0
	for _, s := range teams {
		ratio := float64(s.succ) / float64(s.total)
		if ratio > best {
			best = ratio
		}
	}
	return best
}
```
