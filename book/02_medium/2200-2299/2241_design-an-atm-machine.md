# 2241 — Design An Atm Machine

## Deskripsi

**Soal:** [2241. Design An Atm Machine](https://leetcode.com/problems/design-an-atm-machine/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor() ATM`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
