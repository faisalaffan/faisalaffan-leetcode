package main

// LeetCode #3037: Find Pattern in Infinite Stream II
// https://leetcode.com/problems/find-pattern-in-infinite-stream-ii/
// Difficulty: Hard [Paid]
//
// Given an infinite binary stream that repeats a given bit array cyclically,
// find the first occurrence index of a given pattern.
//
// Approach: KMP algorithm on an infinite stream
//   Build the LPS array for the pattern. Read bits one by one from the stream,
//   running KMP matching. Since the stream is infinite but periodic, we stop
//   after at most len(stream)+len(pattern) reads (because after that, if not found,
//   the pattern cannot occur starting within the periodic prefix; but in an
//   infinite cyclically repeating stream, the pattern must occur at some position
//   within the first cycle + pattern length).

import "fmt"

type InfiniteStream struct {
	bits []int
	pos  int
}

func NewInfiniteStream(bits []int) *InfiniteStream {
	return &InfiniteStream{bits: bits, pos: 0}
}

func (s *InfiniteStream) Next() int {
	val := s.bits[s.pos%len(s.bits)]
	s.pos++
	return val
}

func findPattern(stream *InfiniteStream, pattern []int) int {
	m := len(pattern)

	// Build LPS array for the pattern
	lps := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		lps[i] = j
	}

	// Search in the infinite stream
	// We need to bound the search. The stream is periodic with period = len(bits).
	// The pattern can start at any index in the cyclic stream.
	// We read up to the point where we've covered one full cycle + pattern length.
	// If the pattern exists, it will be found within the first len(stream)+len(pattern)-1
	// positions of the repeated stream.
	j = 0
	limit := stream.pos + len(stream.bits) + m
	for idx := 0; ; idx++ {
		bit := stream.Next()
		for j > 0 && bit != pattern[j] {
			j = lps[j-1]
		}
		if bit == pattern[j] {
			j++
		}
		if j == m {
			return idx - m + 1
		}
		// Safety: break after sufficient reads to avoid infinite loop
		// In practice for a repeating stream, the pattern must be found;
		// but we add a bound for safety.
		if stream.pos > limit+m*2 {
			break
		}
	}
	return -1
}

func main() {
	// Test 1: Simple case
	stream := NewInfiniteStream([]int{1, 1, 1, 0, 1, 1, 1})
	fmt.Println("Test 1:", findPattern(stream, []int{0, 1}))

	// Test 2: Pattern at the beginning
	stream2 := NewInfiniteStream([]int{1, 0, 1, 0, 1})
	fmt.Println("Test 2:", findPattern(stream2, []int{1, 0}))

	// Test 3: Pattern wraps around (occurs across cycle boundary)
	stream3 := NewInfiniteStream([]int{0, 1, 1, 0})
	fmt.Println("Test 3:", findPattern(stream3, []int{0, 0})) // wraps: ...0[0,1,1,0,0,1,1,0]...

	// Test 4: All ones pattern
	stream4 := NewInfiniteStream([]int{1, 0, 1})
	fmt.Println("Test 4:", findPattern(stream4, []int{1, 1}))

	// Test 5: Single element pattern
	stream5 := NewInfiniteStream([]int{0, 1, 1, 0})
	fmt.Println("Test 5:", findPattern(stream5, []int{1}))
}
