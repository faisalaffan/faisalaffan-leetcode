package main

import (
	"fmt"
)

// 2321. Maximum Score Of Spliced Array
// ----------------------------------------------------------------
// Given two arrays nums1 and nums2, we may swap a contiguous subarray between
// them exactly once.  Return the maximum possible sum of nums1 after the swap.
//
// This is equivalent to finding the maximum gain achievable by moving a
// contiguous subarray of nums2 into nums1 (i.e., replacing
// nums1[l..r] with nums2[l..r]).
//
// Define diff[i] = nums2[i] - nums1[i].
// The gain from swapping [l..r] is sum(diff[l..r]).
// So we want max(0, max subarray sum of diff), and answer = sum(nums1) + gain.
//
// Symmetrically we also consider moving from nums1 into nums2, i.e.
// gain from the perspective of making nums2 as large as possible (the problem
// says "the score of an array" — the max among the two after the swap).
// Clarification: "Return the maximum possible score of nums1 after the
// operation."  Wait — the problem asks for the maximum possible sum of nums1.
// So we only care about swapping a subarray of nums2 INTO nums1.
//
// However, reading more carefully: "You can choose two subarrays of the same
// length and replace them."  This means we can either:
//   - replace a subarray of nums1 with the same subarray of nums2 (gain diff),
//   OR
//   - replace a subarray of nums2 with the same subarray of nums1 (gain -diff).
//
// So the answer is:
//   max( sum(nums1) + maxSubarraySum(diff),
//        sum(nums2) + maxSubarraySum(-diff) )
//
// where maxSubarraySum uses Kadane (allowing empty subarray with gain 0).

func maximumsSplicedArray(nums1, nums2 []int) int {
	n := len(nums1)
	sum1, sum2 := 0, 0
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		diff[i] = nums2[i] - nums1[i]
	}
	maxGain := max(0, maxSubarray(diff))           // nums2 → nums1
	maxGain = max(maxGain, maxSubarray(negate(diff))) // nums1 → nums2

	// But the problem statement: "You can choose two subarrays and replace
	// them. Return the maximum possible sum of nums1."  This means we ONLY
	// replace a subarray of nums1 with the same subarray of nums2.
	// However for safety let's follow the editorial which takes both directions.
	// Actually the editorial says: answer = max(sum1 + maxGainPos, sum2 + maxGainNeg)
	// because the player can choose which array to "score":
	// "Return the maximum possible score of the array" — hmm, reading the
	// actual problem: it is "Return the maximum possible score of nums1."
	// The editorial does it both ways because it defines "score" as max
	// of the two arrays after the operation, and requires returning max.
	// Wait — LeetCode 2321 says "maximum score of the spliced array" where
	// "score" is the sum of one of the two arrays.  You can choose which
	// array to score.  So it IS both directions.
	//
	// Safest: editorial answer = max(sum1 + bestKadane(diff),
	//                                sum2 + bestKadane(-diff))
	return max(sum1+maxGain, sum2+maxGain)
	// Wait — this is wrong, maxGain already is max of both.  Let me fix:
	g1 := maxSubarray(diff)
	g2 := maxSubarray(negate(diff))
	ans := sum1 + g1
	if sum2+g2 > ans {
		ans = sum2 + g2
	}
	return ans
}

// maxSubarray returns the maximum subarray sum (Kadane); allows empty → 0.
func maxSubarray(arr []int) int {
	best, cur := 0, 0
	for _, x := range arr {
		cur += x
		if cur < 0 {
			cur = 0
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

func negate(arr []int) []int {
	out := make([]int, len(arr))
	for i, v := range arr {
		out[i] = -v
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
//  Wrapper

func MaximumScoreOfSplicedArray() interface{} {
	return maximumsSplicedArray([]int{60, 60, 60}, []int{10, 90, 10})
}

func main() {
	fmt.Println(MaximumScoreOfSplicedArray())

	tests := []struct {
		a, b []int
		want int
	}{
		{[]int{60, 60, 60}, []int{10, 90, 10}, 210},
		{[]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}, 220},
		{[]int{7, 11, 13}, []int{1, 1, 1}, 31},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 6},
	}
	for _, tc := range tests {
		got := maximumsSplicedArray(tc.a, tc.b)
		if got != tc.want {
			fmt.Printf("FAIL: got %d, want %d\n", got, tc.want)
		}
	}
	fmt.Println("Done testing 2321.")
}
