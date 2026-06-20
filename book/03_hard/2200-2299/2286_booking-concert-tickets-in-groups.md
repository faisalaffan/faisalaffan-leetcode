# 2286 — Booking Concert Tickets In Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n, m int) BookMyShow
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Segment Tree

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import "fmt"

// 2286. Booking Concert Tickets in Groups
// ----------------------------------------------------------------
// Maintain n rows each with m seats.  Two operations:
//
//   gather(k, maxRow):
//     Find the first row ≤ maxRow with ≥ k consecutive free seats.
//     Book those k seats in that row, returning [row, col].
//     If impossible return [nil] (empty slice).
//
//   scatter(k, maxRow):
//     Book k seats among rows [0 … maxRow] greedily from left.
//     Return true on success, false otherwise.
//
// Both require O(log n) each — hence a segment tree storing per node:
//   sumFree  – total free seats in the interval
//   maxConsec – maximum consecutive free seats in any row of the interval
//              (this is just max of per‑row free seats because seats in
//               different rows are not consecutive per the problem).

type segNode struct {
	sumFree   int
	maxConsec int
}

type BookMyShow struct {
	m        int
	n        int
	free     []int   // free[i] = free seats remaining in row i
	seg      []segNode
}

func Constructor(n, m int) BookMyShow {
	size := 4 * n
  // Alokasi slice integer
	free := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range free {
		free[i] = m
	}
	seg := make([]segNode, size)
	b := BookMyShow{m: m, n: n, free: free, seg: seg}
	b.build(1, 0, n-1)
	return b
}

func (b *BookMyShow) build(idx, l, r int) {
	if l == r {
		b.seg[idx] = segNode{sumFree: b.m, maxConsec: b.m}
		return
	}
	mid := (l + r) / 2
	b.build(idx*2, l, mid)
	b.build(idx*2+1, mid+1, r)
	b.pull(idx)
}

func (b *BookMyShow) pull(idx int) {
	b.seg[idx].sumFree = b.seg[idx*2].sumFree + b.seg[idx*2+1].sumFree
	if b.seg[idx*2].maxConsec > b.seg[idx*2+1].maxConsec {
		b.seg[idx].maxConsec = b.seg[idx*2].maxConsec
	} else {
		b.seg[idx].maxConsec = b.seg[idx*2+1].maxConsec
	}
}

// Gather finds the first row ≤ maxRow with ≥ k free seats,
// books them, and returns [row, col].  Returns empty slice on failure.
func (b *BookMyShow) Gather(k int, maxRow int) []int {
	if k <= 0 || maxRow < 0 || b.seg[1].maxConsec < k {
		return []int{}
	}
	// Binary search on segment tree for leftmost row ≤ maxRow with ≥ k free.
	row := b.findFirst(1, 0, b.n-1, maxRow, k)
	if row == -1 {
		return []int{}
	}
	// Book k seats starting at column (b.M - free[row]).
	col := b.m - b.free[row]
	b.free[row] -= k
	b.update(1, 0, b.n-1, row, b.free[row])
	return []int{row, col}
}

// findFirst returns the smallest row ∈ [0, limit] with free ≥ need, or -1.
func (b *BookMyShow) findFirst(idx, l, r, limit, need int) int {
	if l > limit || b.seg[idx].maxConsec < need {
		return -1
	}
	if l == r {
		if b.free[l] >= need {
			return l
		}
		return -1
	}
	mid := (l + r) / 2
	res := b.findFirst(idx*2, l, mid, limit, need)
	if res != -1 {
		return res
	}
	return b.findFirst(idx*2+1, mid+1, r, limit, need)
}

// Scatter attempts to book k seats across rows [0 … maxRow].
func (b *BookMyShow) Scatter(k int, maxRow int) bool {
	if k <= 0 || maxRow < 0 {
		return false
	}
	if b.querySum(1, 0, b.n-1, 0, maxRow) < k {
		return false
	}
	remaining := k
	// Walk rows from 0 upward, taking whatever is free.
	for r := 0; r <= maxRow && remaining > 0; r++ {
		if b.free[r] == 0 {
			continue
		}
		take := remaining
		if b.free[r] < take {
			take = b.free[r]
		}
		b.free[r] -= take
		b.update(1, 0, b.n-1, r, b.free[r])
		remaining -= take
	}
	return true
}

func (b *BookMyShow) querySum(idx, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		return b.seg[idx].sumFree
	}
	mid := (l + r) / 2
	return b.querySum(idx*2, l, mid, ql, qr) + b.querySum(idx*2+1, mid+1, r, ql, qr)
}

func (b *BookMyShow) update(idx, l, r, pos, val int) {
	if l == r {
		b.seg[idx] = segNode{sumFree: val, maxConsec: val}
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		b.update(idx*2, l, mid, pos, val)
	} else {
		b.update(idx*2+1, mid+1, r, pos, val)
	}
	b.pull(idx)
}

// ---------------------------------------------------------------------------
//  Wrapper

func BookingConcertTicketsInGroups() interface{} {
	show := Constructor(2, 5)
  // Alokasi slice integer
	out := make([]interface{}, 0)
	out = append(out, show.Scatter(4, 0)) // true
	out = append(out, show.Scatter(2, 0)) // true
	out = append(out, show.Gather(5, 1))  // [0,0] (row0 had 1 seat left after scatter)
	return out
}

func main() {
	fmt.Println(BookingConcertTicketsInGroups())

	// ---- test ----
	show := Constructor(5, 10)
	if got := show.Gather(6, 3); len(got) == 0 || got[0] != 0 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(6,3) expected [0,0], got %v\n", got)
	}
	if got := show.Gather(5, 2); len(got) == 0 || got[0] != 1 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(5,2) expected [1,0], got %v\n", got)
	}
	if got := show.Gather(12, 4); len(got) != 0 {
		fmt.Printf("FAIL: Gather(12,4) expected [], got %v\n", got)
	}
	if got := show.Scatter(30, 4); got != true {
		fmt.Printf("FAIL: Scatter(30,4) expected true, got %v\n", got)
	}
	if got := show.Gather(2, 4); len(got) == 0 || got[0] != 4 || got[1] != 1 {
		fmt.Printf("FAIL: Gather(2,4) expected [4,1], got %v\n", got)
	}

	show2 := Constructor(3, 3)
	if got := show2.Scatter(2, 2); got != true {
		fmt.Printf("FAIL: Scatter(2,2) expected true\n")
	}
	if got := show2.Gather(3, 2); len(got) == 0 || got[0] != 1 || got[1] != 0 {
		fmt.Printf("FAIL: Gather(3,2) expected [1,0], got %v\n", got)
	}
	fmt.Println("Done testing 2286.")
}
```
