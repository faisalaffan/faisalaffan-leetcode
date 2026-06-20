# 0810 — Chalkboard Xor Game

## Deskripsi

**Soal:** [0810. Chalkboard Xor Game](https://leetcode.com/problems/chalkboard-xor-game/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #810: Chalkboard XOR Game
// https://leetcode.com/problems/chalkboard-xor-game/
// Difficulty: Hard
//
// Two players alternately remove one number from an array. If after a move the
// XOR of all remaining numbers is 0, that player wins.
//
// Alice goes first. Determine if she can force a win assuming optimal play.
//
// Key insight:
//   - If XOR of all numbers is 0, Alice wins immediately.
//   - If n is even, Alice can always win by parity argument: with even length,
//     Alice makes the last move; after she removes the final number, XOR of
//     the empty set is 0.
//   - If n is odd, Bob wins.
//
// Proof: When XOR != 0 and length is even, there is always at least one number
// Alice can remove that leaves a non-zero XOR, so Bob never sees XOR=0. Since
// Bob removes from an odd-length array but Alice started with even, Alice
// removes the last element and wins.

import "fmt"

func main() {
	// Example: [1,1,2] → false (n=3 odd, XOR=2 != 0)
	fmt.Println(xorGame([]int{1, 1, 2}))
	// [1,1] → true (n=2 even)
	fmt.Println(xorGame([]int{1, 1}))
	// [1,2,3] → true? n=3 odd but XOR=0 (1^2^3=0) → Alice wins immediately
	fmt.Println(xorGame([]int{1, 2, 3}))
	// [0] → true (XOR=0)
	fmt.Println(xorGame([]int{0}))
	// [1] → false (n=1 odd, XOR=1 != 0)
	fmt.Println(xorGame([]int{1}))
}

func xorGame(nums []int) bool {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	return xor == 0 || len(nums)%2 == 0
}
```
