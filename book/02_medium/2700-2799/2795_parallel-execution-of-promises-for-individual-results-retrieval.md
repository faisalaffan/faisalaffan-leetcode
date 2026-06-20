# 2795 — Parallel Execution Of Promises For Individual Results Retrieval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ParallelExecutionOfPromisesForIndividualResultsRetrieval(functions []func() int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2795: Parallel Execution of Promises for Individual Results Retrieval
// https://leetcode.com/problems/parallel-execution-of-promises-for-individual-results-retrieval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

type PromiseResult struct {
	Index int
	Value interface{}
}

func ParallelExecutionOfPromisesForIndividualResultsRetrieval(functions []func() int) []int {
	n := len(functions)
  // Alokasi slice integer
	results := make([]int, n)
	var wg sync.WaitGroup

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 1 + 1 },
		func() int { return 2 + 2 },
		func() int { return 3 + 3 },
	})
	fmt.Println(results)

	results2 := ParallelExecutionOfPromisesForIndividualResultsRetrieval([]func() int{
		func() int { return 99 },
	})
	fmt.Println(results2)
}
```
