# 3490 — Count Beautiful Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func init() 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3490: Count Beautiful Numbers
// https://leetcode.com/problems/count-beautiful-numbers/
// Difficulty: Hard
//
// A number is "beautiful" if all digits are non-zero and the product of
// its digits divides the number.
//
// Digit DP tracking remainders modulo prime powers (2^a, 3^b, 5^c, 7^d).
// At the leaf, check num % (2^a * 3^b * 5^c * 7^d) == 0 by checking
// each prime power condition independently: rem % 2^a == 0 etc.

import "fmt"

const (
	MAX2 = 30
	MAX3 = 18
	MAX5 = 9
	MAX7 = 9
)

var (
	pow2 [MAX2 + 1]int
	pow3 [MAX3 + 1]int
	pow5 [MAX5 + 1]int
	pow7 [MAX7 + 1]int
)

func init() {
	pow2[0], pow3[0], pow5[0], pow7[0] = 1, 1, 1, 1
	for i := 1; i <= MAX2; i++ {
		pow2[i] = pow2[i-1] * 2
	}
	for i := 1; i <= MAX3; i++ {
		pow3[i] = pow3[i-1] * 3
	}
	for i := 1; i <= MAX5; i++ {
		pow5[i] = pow5[i-1] * 5
	}
	for i := 1; i <= MAX7; i++ {
		pow7[i] = pow7[i-1] * 7
	}
}

// Prime factor exponent contributions for digits 0-9
// digit -> (e2, e3, e5, e7)
var expTable = [10][4]int{
	{0, 0, 0, 0}, // 0
	{0, 0, 0, 0}, // 1
	{1, 0, 0, 0}, // 2
	{0, 1, 0, 0}, // 3
	{2, 0, 0, 0}, // 4
	{0, 0, 1, 0}, // 5
	{1, 1, 0, 0}, // 6
	{0, 0, 0, 1}, // 7
	{3, 0, 0, 0}, // 8
	{0, 2, 0, 0}, // 9
}

func countBeautifulNumbers(low, high int) int {
	if low < 1 {
		low = 1
	}
	return countUpTo(high) - countUpTo(low-1)
}

func countUpTo(limit int) int {
	if limit <= 0 {
		return 0
	}
	// Extract digits
	var digs [20]int
	n := 0
	for tmp := limit; tmp > 0; tmp /= 10 {
		digs[n] = tmp % 10
		n++
	}
	// Reverse
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		digs[i], digs[j] = digs[j], digs[i]
	}

	type stateKey struct {
		a, b, c, d byte
		r2, r3     int
		r5, r7     int
	}
  // Membuat matriks/slice 2D untuk DP
	memo := make([][][]map[stateKey]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([][]map[stateKey]int, 2)
		memo[i][0] = make([]map[stateKey]int, 2)
		memo[i][1] = make([]map[stateKey]int, 2)
	}

	var dfs func(pos, tight, started int, a, b, c, d int, r2, r3, r5, r7 int) int
	dfs = func(pos, tight, started int, a, b, c, d int, r2, r3, r5, r7 int) int {
		if pos == n {
			if started == 0 {
				return 0
			}
			if r2%pow2[a] == 0 && r3%pow3[b] == 0 && r5%pow5[c] == 0 && r7%pow7[d] == 0 {
				return 1
			}
			return 0
		}

		sk := stateKey{byte(a), byte(b), byte(c), byte(d), r2, r3, r5, r7}
		if tight == 0 {
			if memo[pos][tight][started] == nil {
				memo[pos][tight][started] = make(map[stateKey]int)
			} else if val, ok := memo[pos][tight][started][sk]; ok {
				return val
			}
		}

		maxD := 9
		if tight == 1 {
			maxD = digs[pos]
		}

		total := 0
		for dg := 0; dg <= maxD; dg++ {
			nt := 0
			if tight == 1 && dg == maxD {
				nt = 1
			}
			if started == 0 && dg == 0 {
				total += dfs(pos+1, nt, 0, 0, 0, 0, 0, 0, 0, 0, 0)
			} else if dg == 0 {
				// Digit 0 → product 0, cannot divide
				continue
			} else {
				e := expTable[dg]
				na := a + e[0]
				nb := b + e[1]
				nc := c + e[2]
				nd := d + e[3]
				if na > MAX2 || nb > MAX3 || nc > MAX5 || nd > MAX7 {
					continue
				}
				nr2 := (r2*10 + dg) % pow2[MAX2]
				nr3 := (r3*10 + dg) % pow3[MAX3]
				nr5 := (r5*10 + dg) % pow5[MAX5]
				nr7 := (r7*10 + dg) % pow7[MAX7]
				total += dfs(pos+1, nt, 1, na, nb, nc, nd, nr2, nr3, nr5, nr7)
			}
		}

		if tight == 0 {
			if memo[pos][tight][started] == nil {
				memo[pos][tight][started] = make(map[stateKey]int)
			}
			memo[pos][tight][started][sk] = total
		}
		return total
	}

	return dfs(0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0)
}

func main() {
	fmt.Printf("count(1,20) -> %d\n", countBeautifulNumbers(1, 20))
	fmt.Printf("count(1,100) -> %d\n", countBeautifulNumbers(1, 100))
	fmt.Printf("count(1,1000) -> %d\n", countBeautifulNumbers(1, 1000))
	fmt.Printf("count(1,10000) -> %d\n", countBeautifulNumbers(1, 10000))
}
```
