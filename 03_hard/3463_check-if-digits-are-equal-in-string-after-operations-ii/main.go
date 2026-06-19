package main

// LeetCode #3463: Check If Digits Are Equal in String After Operations II
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-ii/
// Difficulty: Hard
//
// Repeatedly replace adjacent pair (a,b) with (a+b)%10 until 2 digits remain.
// Final two digits are equal iff:
//   sum_{j=0}^{n-2} C(n-2, j) * int(s[j])   ≡
//   sum_{j=0}^{n-2} C(n-2, j) * int(s[j+1]) (mod 10)
//
// Compute C(n,k) mod 10 via Lucas theorem mod 2 and mod 5, then CRT.
// n up to 10^5 so factorial precomputation is fine.

import "fmt"

// factorials modulo 5
var fact5 = [5]int{1, 1, 2, 6, 24} // 0!,1!,2!,3!,4! values
var invFact5 = [5]int{1, 1, 3, 2, 4} // modular inverses: fact5[i] * invFact5[i] ≡ 1 (mod 5)

// Cmod5(n,k): binomial coefficient modulo 5 using Lucas theorem
func binomMod5(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if n < 5 {
		return fact5[n] / (fact5[k] * fact5[n-k]) % 5
	}
	// Lucas: represent n,k in base 5, multiply C(base5_digit_n, base5_digit_k) mod 5
	res := 1
	for n > 0 || k > 0 {
		ni := n % 5
		ki := k % 5
		if ki > ni {
			return 0
		}
		num := fact5[ni]
		den := fact5[ki] * fact5[ni-ki] % 5
		// modular inverse of den mod 5
		invDen := invFact5[den]
		res = res * num % 5 * invDen % 5
		n /= 5
		k /= 5
	}
	return res
}

// binomMod2: binomial coefficient modulo 2 using Lucas theorem
// C(n,k) mod 2 = 1 iff (k & ~n) == 0 i.e. k is a submask of n
func binomMod2(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k&^n == 0 {
		return 1
	}
	return 0
}

// binomMod10: binomial coefficient modulo 10 via CRT
func binomMod10(n, k int) int {
	r2 := binomMod2(n, k)
	r5 := binomMod5(n, k)
	// Solve x ≡ r2 (mod 2), x ≡ r5 (mod 5)
	// x = r5 * 6 + r2 * 5 (mod 10)
	return (r5*6 + r2*5) % 10
}

func isEqualAfterOps(s string) bool {
	n := len(s)
	if n <= 2 {
		return s[0] == s[1]
	}

	m := n - 2
	sum0 := 0
	sum1 := 0
	for j := 0; j < n-1; j++ {
		c := binomMod10(m, j)
		if j < n {
			sum0 = (sum0 + c*int(s[j]-'0')) % 10
		}
		if j+1 < n {
			sum1 = (sum1 + c*int(s[j+1]-'0')) % 10
		}
	}
	return sum0 == sum1
}

func main() {
	// Test: "3902" -> expected true
	fmt.Printf("s=3902 -> %v (expected true)\n", isEqualAfterOps("3902"))

	// Test: "12" -> expected true
	fmt.Printf("s=12 -> %v (expected true)\n", isEqualAfterOps("12"))

	// Test: "3478" -> simulate: 3478->717->88, so true
	fmt.Printf("s=3478 -> %v (expected true)\n", isEqualAfterOps("3478"))

	// Test: "0000" -> expected true
	fmt.Printf("s=0000 -> %v (expected true)\n", isEqualAfterOps("0000"))

	// Test: "1234" -> expected false (simulated above -> 82)
	fmt.Printf("s=1234 -> %v (expected false)\n", isEqualAfterOps("1234"))

	// Test: "11" -> expected true
	fmt.Printf("s=11 -> %v (expected true)\n", isEqualAfterOps("11"))
}
