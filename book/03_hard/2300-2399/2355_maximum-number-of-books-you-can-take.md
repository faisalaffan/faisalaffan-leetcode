# 2355 — Maximum Number Of Books You Can Take

## Deskripsi

**Soal:** [2355. Maximum Number Of Books You Can Take](https://leetcode.com/problems/maximum-number-of-books-you-can-take/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Stack (tumpukan LIFO), Trie (pohon awalan)

**Fungsi Solusi:** `func maximumBooks(books []int) int64`

## Solusi Go

```go
package main

import (
	"fmt"
)

// 2355. Maximum Number of Books You Can Take
// ----------------------------------------------------------------
// Choose a contiguous segment of shelves.  Within the segment you must take
// a non‑increasing number of books (a[i] ≥ a[i+1]) and at most books[i] from
// each shelf.  Maximise total books taken.
//
// For segment [l, r]: a[l]=books[l], a[t]=min(books[t], a[t-1]) = running min.
//
// DP with monotonic stack.  Each entry: (minVal, maxSum).
// At position i:
//   - Pop entries with minVal ≥ books[i]; track best sum among them.
//   - Extend surviving entries: add their own minVal (running min unchanged).
//   - Push new entry for segments ending at i with running min = books[i]:
//     sum = bestMerged + books[i]  (best from extending popped entries).
//   - Answer = max over all entries' sums.

type entry struct {
	minVal int
	sum    int64
}

func maximumBooks(books []int) int64 {
  // Membuat slice untuk menyimpan hasil
	stack := make([]entry, 0)
	var result int64

	for _, bi := range books {
		var bestMerged int64 // best sum among popped entries

		// Pop entries with minVal ≥ bi.
		for len(stack) > 0 && stack[len(stack)-1].minVal >= bi {
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if e.sum > bestMerged {
				bestMerged = e.sum
			}
		}

		// Extend survivors: their running min (minVal) doesn't change.
		for j := range stack {
			stack[j].sum += int64(stack[j].minVal)
		}

		// New entry with running min = bi.
		newSum := bestMerged + int64(bi)
		stack = append(stack, entry{bi, newSum})
		if newSum > result {
			result = newSum
		}

		// Also check best survivor sum.
		if len(stack) > 1 {
			// stack[0..len(stack)-2] are survivors (their sums were just updated).
			for j := 0; j < len(stack)-1; j++ {
				if stack[j].sum > result {
					result = stack[j].sum
				}
			}
		}
	}
	return result
}

// ---------------------------------------------------------------------------
//  Wrapper

func MaximumNumberOfBooksYouCanTake() interface{} {
	return maximumBooks([]int{8, 5, 2, 7, 7})
}

func main() {
	fmt.Println(MaximumNumberOfBooksYouCanTake())

	cases := []struct {
		books []int
		want  int64
	}{
		{[]int{8, 5, 2, 7, 7}, 19},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{5, 5, 5}, 15},
		{[]int{7, 0, 0, 0, 7}, 7},
		{[]int{10, 1, 1, 1, 1, 1}, 15},
		{[]int{3, 0, 5, 0, 2}, 5},
		{[]int{2, 2, 2, 2, 2}, 10},
	}
	for _, c := range cases {
		if got := maximumBooks(c.books); got != c.want {
			fmt.Printf("FAIL books=%v: got %d, want %d\n", c.books, got, c.want)
		}
	}
	fmt.Println("Done testing 2355.")
}
```
