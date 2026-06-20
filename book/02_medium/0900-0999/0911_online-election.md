# 0911 — Online Election

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(persons []int, times []int) TopVotedCandidate
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #911: Online Election
// https://leetcode.com/problems/online-election/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type TopVotedCandidate struct {
	times []int
	wins  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(times)
  // Alokasi slice integer
	wins := make([]int, n)
  // Membuat map (HashMap) — pencarian O(1)
	votes := make(map[int]int)
	leader := -1

	for i := 0; i < n; i++ {
		votes[persons[i]]++
		if leader == -1 || votes[persons[i]] >= votes[leader] {
			leader = persons[i]
		}
		wins[i] = leader
	}

	return TopVotedCandidate{times, wins}
}

func (this *TopVotedCandidate) Q(t int) int {
	idx := sort.SearchInts(this.times, t)
	if idx < len(this.times) && this.times[idx] == t {
		return this.wins[idx]
	}
	return this.wins[idx-1]
}

func main() {
	obj := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(obj.Q(3))
	fmt.Println(obj.Q(12))
	fmt.Println(obj.Q(25))
	fmt.Println(obj.Q(15))
	fmt.Println(obj.Q(24))
	fmt.Println(obj.Q(8))
}
```
