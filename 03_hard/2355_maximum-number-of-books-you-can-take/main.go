package main

import (
	"fmt"
)

// 2355. Maximum Number of Books You Can Take
// ----------------------------------------------------------------
// Given an array books[n] where you may take at most books[i] books from
// shelf i.  You must choose a contiguous segment and, within it, take a
// non‑increasing sequence of books (a[i] >= a[i+1] >= ...).
// Maximise the total number of books taken.
//
// DP with monotonic stack (O(n)):
//   Let dp[i] = max sum for a segment ending at i.
//   Define val[i] = books[i] - i.
//   Maintain a stack of indices with *strictly increasing* val[i] (bottom to
//   top).  For each i, pop while val[top] >= val[i].
//
//   After popping, let j = stack top (or -1 if empty).
//   For the segment from j+1 to i (inclusive), the values form a perfect
//   arithmetic progression: books[i], books[i]-1, ..., books[i]-len+1
//   where len = min(i - j, books[i]).
//
//   Proof: For any p in (j,i], books[p] - p >= val[i] (by the pop condition),
//   so books[p] >= val[i] + p = books[i] - i + p.  The ideal non‑increasing
//   amount at p is books[i] - (i-p) = books[i] - i + p, which is ≤ books[p].
//   Hence the progression is not capped by the shelf capacity, and the
//   non‑increasing constraint is satisfied by construction.
//
//   dp[i] = (j >= 0 ? dp[j] : 0) + arithmeticSum(books[i], len).

func maximumBooks(books []int) int64 {
	n := len(books)
	dp := make([]int64, n)
	stack := make([]int, 0)
	var result int64

	for i := 0; i < n; i++ {
		// Pop while val[top] >= val[i].
		val := books[i] - i
		for len(stack) > 0 && books[stack[len(stack)-1]]-stack[len(stack)-1] >= val {
			stack = stack[:len(stack)-1]
		}

		var j int
		if len(stack) == 0 {
			j = -1
		} else {
			j = stack[len(stack)-1]
		}

		// Arithmetic sum from j+1 to i.
		length := i - j // number of positions in the arithmetic tail
		if length > books[i] {
			length = books[i]
		}
		// sum = length * (last + last - length + 1) / 2
		first := int64(books[i] - length + 1)
		last := int64(books[i])
		sum := (first + last) * int64(length) / 2

		if j >= 0 {
			dp[i] = dp[j] + sum
		} else {
			dp[i] = sum
		}

		if dp[i] > result {
			result = dp[i]
		}
		stack = append(stack, i)
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

	// Test cases.
	testCases := []struct {
		books []int
		want  int64
	}{
		{[]int{8, 5, 2, 7, 7}, 19},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 5, 5}, 15},
		{[]int{7, 0, 0, 0, 7}, 14},
		{[]int{2, 2, 2, 2, 2}, 10},
		{[]int{10, 1, 1, 1, 1, 1}, 15},
	}
	for _, tc := range testCases {
		if got := maximumBooks(tc.books); got != tc.want {
			fmt.Printf("FAIL books=%v: got %d, want %d\n", tc.books, got, tc.want)
		}
	}

	// Additional test.
	got := maximumBooks([]int{3, 0, 5, 0, 2})
	fmt.Printf("maximumBooks([3,0,5,0,2]) = %d (expected 5)\n", got)

	fmt.Println("Done testing 2355.")
}
