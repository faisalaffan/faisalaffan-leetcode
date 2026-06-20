# 2241 — Design An Atm Machine

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() ATM
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2241: Design an ATM Machine
// https://leetcode.com/problems/design-an-atm-machine/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(1)

import "fmt"

type ATM struct {
	denoms []int
	notes  []int
}

func Constructor() ATM {
	return ATM{
		denoms: []int{20, 50, 100, 200, 500},
		notes:  make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, c := range banknotesCount {
		this.notes[i] += c
	}
}

func (this *ATM) Withdraw(amount int) []int {
  // Alokasi slice integer
	result := make([]int, 5)
	need := amount
	for i := 4; i >= 0; i-- {
		use := need / this.denoms[i]
		if use > this.notes[i] {
			use = this.notes[i]
		}
		result[i] = use
		need -= use * this.denoms[i]
	}
	if need != 0 {
		return []int{-1}
	}
	for i, c := range result {
		this.notes[i] -= c
	}
	return result
}

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
	// Expected: [0,0,1,0,1]

	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println(atm.Withdraw(600))
	// Expected: [-1]

	fmt.Println(atm.Withdraw(550))
	// Expected: [0,1,0,0,1]
}
```
