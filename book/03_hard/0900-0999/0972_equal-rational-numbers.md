# 0972 — Equal Rational Numbers

## Deskripsi

**Soal:** [0972. Equal Rational Numbers](https://leetcode.com/problems/equal-rational-numbers/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func isRationalEqual(s string, t string) bool`

## Solusi Go

```go
package main

// LeetCode #972: Equal Rational Numbers
// https://leetcode.com/problems/equal-rational-numbers/
// Difficulty: Hard

import (
	"fmt"
	"math/big"
	"strings"
)

func isRationalEqual(s string, t string) bool {
	return parseRational(s).Cmp(parseRational(t)) == 0
}

// parseRational converts a string like "0.(52)" or "0.5(25)" to a big.Rat
func parseRational(s string) *big.Rat {
	// Split into integer, decimal, and repeating parts
	// Format: int.dec(repeat) or int.dec or int

	hasDot := strings.Contains(s, ".")
	if !hasDot {
		// Just an integer
		val := new(big.Rat)
		val.SetString(s)
		return val
	}

	parts := strings.SplitN(s, ".", 2)
	intPart := parts[0]
	decPart := parts[1]

	// Check for repeating part
	openParen := strings.Index(decPart, "(")
	var nonRepeat, repeat string
	if openParen != -1 {
		nonRepeat = decPart[:openParen]
		repeat = decPart[openParen+1 : len(decPart)-1] // remove '(' and ')'
	} else {
		nonRepeat = decPart
		repeat = ""
	}

	// Build rational number:
	// integer part + non-repeating part / 10^len(nonRepeat)
	// + repeating part / (10^len(nonRepeat) * (10^len(repeat) - 1))

	result := new(big.Rat)

	// Integer part
	intNum := new(big.Int)
	intNum.SetString(intPart, 10)
	intRat := new(big.Rat).SetInt(intNum)
	result.Add(result, intRat)

	if len(nonRepeat) > 0 {
		// nonRepeating / 10^len(nonRepeat)
		num := new(big.Int)
		num.SetString(nonRepeat, 10)
		denom := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(len(nonRepeat))), nil)
		rat := new(big.Rat).SetFrac(num, denom)
		result.Add(result, rat)
	}

	if len(repeat) > 0 {
		// repeating / ((10^len(repeat) - 1) * 10^len(nonRepeat))
		num := new(big.Int)
		num.SetString(repeat, 10)

		powLen := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(len(repeat))), nil)
		denomBase := new(big.Int).Sub(powLen, big.NewInt(1))

		denom := denomBase
		if len(nonRepeat) > 0 {
			powNonRepeat := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(len(nonRepeat))), nil)
			denom = new(big.Int).Mul(denomBase, powNonRepeat)
		}

		rat := new(big.Rat).SetFrac(num, denom)
		result.Add(result, rat)
	}

	return result
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(isRationalEqual("0.(52)", "0.5(25)"))
	// Expected: true

	fmt.Println("Example 2:")
	fmt.Println(isRationalEqual("0.1666(6)", "0.166(66)"))
	// Expected: true

	fmt.Println("Example 3:")
	fmt.Println(isRationalEqual("0.9(9)", "1."))
	// Expected: true

	fmt.Println("Example 4:")
	fmt.Println(isRationalEqual("1", "1.0"))
	// Expected: true

	fmt.Println("Example 5:")
	fmt.Println(isRationalEqual("0.(8)", "0.(9)"))
	// Expected: false
}
```
