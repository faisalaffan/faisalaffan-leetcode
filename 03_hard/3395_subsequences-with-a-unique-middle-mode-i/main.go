package main

// LeetCode #3395: Subsequences with a Unique Middle Mode I
// https://leetcode.com/problems/subsequences-with-a-unique-middle-mode-i/
// Difficulty: Hard
//
// Count subsequences length 5 with unique middle mode (index 2).
// Fix middle index, use prefix/suffix counts, inclusion-exclusion.

import "fmt"

func main() {
	fmt.Println(SubsequencesWithAUniqueMiddleModeI([]int{1, 2, 2, 3, 3, 4}))
}

func SubsequencesWithAUniqueMiddleModeI(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}

	const MOD = 1000000007

	ans := 0
	for mid := 2; mid <= n-3; mid++ {
		left := make(map[int]int)
		right := make(map[int]int)
		for i := 0; i < mid; i++ {
			left[nums[i]]++
		}
		for i := mid + 1; i < n; i++ {
			right[nums[i]]++
		}
		total := 0

		// Choose 2 from left and 2 from right (any values)
		leftPairs := 0
		for _, c := range left {
			if c >= 2 {
				leftPairs += c * (c - 1) / 2
			}
		}
		rightPairs := 0
		for _, c := range right {
			if c >= 2 {
				rightPairs += c * (c - 1) / 2
			}
		}

		// Total 5-element subsequences with nums[mid] at index 2
		// = (choose 2 from left) * (choose 2 from right)
		// But also consider when same value appears multiple times across left/right

		// For each distinct value v that appears in both sides, subtract invalid
		val := nums[mid]
		// Count subsequences where the middle mode is not unique
		// i.e., some other value appears at least 2 times in the 5 elements.
		// This happens when left has >= 2 of v and right has >= 2 of v,
		// or when some other value appears on both sides such that it could
		// become the mode.

		// Count valid subsequences for this middle:
		// Fix middle = nums[mid]. Need exactly 2 from left and 2 from right.
		// A subsequence is invalid if some other value appears 2+ times.

		lv := left[val]
		rv := right[val]

		// Total: choose2 from left * choose2 from right
		totalLeft := (mid * (mid - 1) / 2) % MOD
		totalRight := ((n - mid - 1) * (n - mid - 2) / 2) % MOD
		total = (totalLeft * totalRight) % MOD

		// Subtract invalid: when left side has 2 of val, right side has 2 of val
		// (then val appears 5 times, not unique middle mode)
		if lv >= 2 && rv >= 2 {
			sub := (lv * (lv - 1) / 2) % MOD
			sub2 := (rv * (rv - 1) / 2) % MOD
			total = (total - sub*sub2%MOD + MOD) % MOD
		}

		// Subtract invalid: when some other value x appears 2+ times across left+right
		// and also appears in the middle (which it can't since middle is fixed).
		// Actually the middle is fixed to nums[mid]. For uniqueness, no other value
		// should appear 2+ times. So if any other value x has:
		//   left[x] + right[x] >= 2, that makes it a competitor for "middle mode"
		//   AND we need at least 1 on each side (since we choose 2 from each side)
		// Actually the issue is: if a value x != val appears in left >= 2 OR right >= 2,
		// then the mode might not be unique.
		// But the mode is defined as the value at index 2 of the chosen 5-length subsequence.
		// The middle mode is unique if no other value appears 2+ times.
		// So invalid if:
		//   - Some x != val appears in left >= 2 OR right >= 2 OR (left[x]>=1 && right[x]>=1)

		for x, lc := range left {
			if x == val {
				continue
			}
			rc := right[x]
			if lc >= 2 || rc >= 2 || (lc >= 1 && rc >= 1) {
				// This value x competes for mode.
				// Count subsequences where x can make the mode non-unique.
				// We need to count subsequences where x appears 2+ times.
				// x can appear 2+ in left, 2+ in right, or 1 each side.
				invalid := 0

				// Case: x appears 2+ in left (x contributes at least 2 from left)
				if lc >= 2 {
					// Need 2 more from left (choose the x's) and 2 from right (any)
					cinvalid := (lc * (lc - 1) / 2) % MOD
					cinvalid = cinvalid * totalRight % MOD
					invalid = (invalid + cinvalid) % MOD
				}

				// Case: x appears 2+ in right
				if rc >= 2 {
					cinvalid := (rc * (rc - 1) / 2) % MOD
					cinvalid = cinvalid * totalLeft % MOD
					invalid = (invalid + cinvalid) % MOD
				}

				// Case: x appears 1 on each side
				if lc >= 1 && rc >= 1 {
					// Need 1 more from left (any of the remaining) and 1 more from right (any of remaining)
					leftRest := (mid - 1)  // left elements excluding the one x
					rightRest := (n - mid - 2) // right elements excluding the one x
					if leftRest >= 1 && rightRest >= 1 {
						cinvalid := lc * rc % MOD
						// After fixing those 2 x's, choose 1 more from left, 1 more from right
						cinvalid = cinvalid * leftRest % MOD
						cinvalid = cinvalid * rightRest % MOD
						invalid = (invalid + cinvalid) % MOD
					}
				}

				total = (total - invalid + MOD) % MOD
			}
		}

		ans = (ans + total) % MOD
	}
	return ans
}
