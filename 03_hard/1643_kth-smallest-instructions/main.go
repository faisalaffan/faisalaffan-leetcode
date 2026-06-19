package main

// LeetCode #1643: Kth Smallest Instructions
// https://leetcode.com/problems/kth-smallest-instructions/
// Difficulty: Hard

import "fmt"

// nCr computes binomial coefficient using Pascal's triangle
func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	res := 1
	for i := 0; i < r; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func kthSmallestPath(destination []int, k int) string {
	v, h := destination[0], destination[1] // vertical (V) and horizontal (H) steps
	n := v + h                              // total steps

	result := make([]byte, n)

	for i := 0; i < n; i++ {
		if h > 0 {
			// Number of ways if we place 'H' here:
			// Remaining positions = v + h - 1, remaining V's = v
			ways := nCr(v+h-1, v)
			if k <= ways {
				result[i] = 'H'
				h--
			} else {
				result[i] = 'V'
				k -= ways
				v--
			}
		} else {
			// No H's left, must place V
			result[i] = 'V'
			v--
		}
	}

	return string(result)
}

func main() {
	// Test case 1: destination=[2,3], k=1 -> "HHHVV"
	dest := []int{2, 3}
	k := 1
	result := kthSmallestPath(dest, k)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHHVV\")\n", dest, k, result)

	// Test case 2: destination=[2,3], k=2 -> "HHVHV"
	k2 := 2
	result2 := kthSmallestPath(dest, k2)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHVHV\")\n", dest, k2, result2)

	// Test case 3: destination=[2,3], k=3 -> "HHVVH"
	k3 := 3
	result3 := kthSmallestPath(dest, k3)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHVVH\")\n", dest, k3, result3)

	// Test case 4: destination=[2,3], k=last -> "VVHHH" (last = C(5,2)=10)
	k4 := 10
	result4 := kthSmallestPath(dest, k4)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"VVHHH\")\n", dest, k4, result4)

	// Test case 5: destination=[1,1], k=1 -> "HV"
	dest5 := []int{1, 1}
	result5 := kthSmallestPath(dest5, 1)
	fmt.Printf("destination=%v k=1 -> \"%s\" (expected \"HV\")\n", dest5, result5)

	// Test case 6: destination=[1,1], k=2 -> "VH"
	result6 := kthSmallestPath(dest5, 2)
	fmt.Printf("destination=%v k=2 -> \"%s\" (expected \"VH\")\n", dest5, result6)
}
