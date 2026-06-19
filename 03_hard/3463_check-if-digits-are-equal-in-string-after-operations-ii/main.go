package main

// LeetCode #3463: Check If Digits Are Equal in String After Operations II
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-ii/
// Difficulty: Hard
//
// Repeatedly replace each adjacent pair (a,b) with (a+b)%10 until 2 digits remain.
// Check if final two digits are equal.
// Uses Lucas theorem to compute binomial coefficients modulo small primes
// (2 and 5) and Chinese Remainder Theorem.

import "fmt"

const MOD = 10

// factor out powers of 2 and 5 from n!, return (remaining_value_mod, count_of_2, count_of_5)
func factor25(n int) (int, int, int) {
	// Compute n! with factors of 2 and 5 removed
	// Use the formula: n! = 2^a * 5^b * r where r is coprime with 10
	val := 1
	cnt2 := 0
	cnt5 := 0

	// We need n! mod something, but excluding factors 2 and 5
	// Use the recursive formula: f(n) = f(n/2) * f(n/5) * product of numbers 1..n not divisible by 2 or 5
	return val, cnt2, cnt5 // placeholder
}

func isEqualAfterOps(s string) bool {
	n := len(s)
	if n <= 2 {
		return s[0] == s[1]
	}

	// After n-2 steps, final digit at position i (0-indexed, where i goes 0..1) is:
	// sum_{j=0}^{n-2} C(n-2, j) * s[j+i] % 10
	// We need to check if these two sums are equal mod 10.

	// Compute C(n-2, k) modulo 10 for all k
	// n up to 10^5 or more, use Lucas theorem mod 2 and mod 5, then CRT

	// Precompute factorials mod 2 and mod 5 using Lucas
	// Lucas theorem: C(n,k) mod p = product C(n_i, k_i) mod p where n_i,k_i are base-p digits

	binomialMod2 := func(n, k int) int {
		// C(n,k) mod 2 = 1 iff (k & ~n) == 0 (Lucas for p=2)
		if k&^n == 0 {
			return 1
		}
		return 0
	}

	// Precompute factorials mod 5 up to 4
	fact5 := [5]int{1, 1, 2, 6, 24}

	var lucasMod5 func(n, k int) int
	lucasMod5 = func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		if n < 5 {
			return fact5[n] / (fact5[k] * fact5[n-k]) % 5
		}
		// Lucas: C(n,k) = C(n/5, k/5) * C(n%5, k%5) mod 5
		// But 5 is prime, n/5 uses integer division
		return lucasMod5(n/5, k/5) * (fact5[n%5] / (fact5[k%5] * fact5[(n-k)%5])) % 5
	}

	// Actually for mod 5 we need modular inverse since 5 is prime
	// Let me use a simpler approach: precompute binomial up to n using
	// formula with 2 and 5 factors removed, then combine with CRT

	// Alternative: compute binomial mod 2 and mod 5 separately

	// Helper: C(n,k) mod 2 using Lucas
	c2 := func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		if k&^n == 0 {
			return 1
		}
		return 0
	}

	// Helper: C(n,k) mod 5 using Lucas
	var c5 func(n, k int) int
	c5 = func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		if n == 0 || k == 0 {
			return 1
		}
		ni := n % 5
		ki := k % 5
		if ki > ni {
			return 0
		}
		// C(ni, ki) mod 5 = fact5[ni] * inv(fact5[ki] * fact5[ni-ki]) mod 5
		num := fact5[ni]
		den := fact5[ki] * fact5[ni-ki] % 5
		// inv of den mod 5 (den is 1,2,3,4 → den^(5-2) = den^3 mod 5)
		invDen := 1
		for d := den; d < 5; d++ {
			if (den*d)%5 == 1 {
				invDen = d
				break
			}
		}
		return num * invDen % 5 * c5(n/5, k/5) % 5
	}

	// CRT: given x ≡ r2 (mod 2) and x ≡ r5 (mod 5), find x mod 10
	// Since 2 and 5 are coprime, solution exists.
	// x = r5 * 6 + r2 * 5 (mod 10) because:
	// 6 ≡ 1 (mod 5), 6 ≡ 0 (mod 2); 5 ≡ 1 (mod 2), 5 ≡ 0 (mod 5)
	crt := func(r2, r5 int) int {
		return (r5*6 + r2*5) % 10
	}

	binomMod10 := func(n, k int) int {
		return crt(c2(n, k), c5(n, k))
	}

	sum0 := 0
	sum1 := 0
	nn := n - 2
	for j := 0; j < n-1; j++ {
		b := binomMod10(nn, j)
		if j < n {
			sum0 = (sum0 + b*int(s[j]-'0')) % 10
		}
		if j+1 < n {
			sum1 = (sum1 + b*int(s[j+1]-'0')) % 10
		}
	}

	return sum0 == sum1
}

func main() {
	// Test: "3902" -> expected true
	fmt.Printf("s=3902 -> %v (expected true)\n", isEqualAfterOps("3902"))

	// Test: "12" -> expected true
	fmt.Printf("s=12 -> %v (expected true)\n", isEqualAfterOps("12"))

	// Test: "3478" -> expected ?
	fmt.Printf("s=3478 -> %v\n", isEqualAfterOps("3478"))

	// Test: "0000" -> expected true
	fmt.Printf("s=0000 -> %v (expected true)\n", isEqualAfterOps("0000"))

	// Test: "1234" -> expected ?
	fmt.Printf("s=1234 -> %v\n", isEqualAfterOps("1234"))
}
