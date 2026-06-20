# 0591 — Tag Validator

## Deskripsi

**Soal:** [0591. Tag Validator](https://leetcode.com/problems/tag-validator/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #591: Tag Validator
// https://leetcode.com/problems/tag-validator/
// Difficulty: Hard

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Test cases
	testCases := []struct {
		input string
		want  bool
	}{
		{"<DIV>This is the first line <![CDATA[<div>]]></DIV>", true},
		{"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>", true},
		{"<A>  <B> </A>   </B>", false},
		{"<DIV>  div tag is not closed  <DIV>", false},
		{"<DIV>  unmatched <  </DIV>", false},
		{"<DIV> closed tags with invalid tag name <b>123</b> </DIV>", false},
		{"<TAG>some text</TAG>", true},
		{"<A></A><B></B>", false},
		{"<![CDATA[wahaha]]>", false},
		{"<A></A>", true},
		{"<A>  </A>", true},
		{"<A></A>>", false},
		{"<A> <B> </B> </A>", true},
		{"<A> <B> </A> </B>", false},
		{"<DIV>  </DIV>", true},
		{"<DIV>  <![CDATA[<div>]]>  </DIV>", true},
		{"<A><![CDATA[<B>]]></A>", true},
	}

	for _, tc := range testCases {
		got := isValid(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: isValid(%q) = %v (want %v)\n", status, tc.input, got, tc.want)
	}
}

func isValid(code string) bool {
	if len(code) == 0 || code[0] != '<' {
		return false
	}

	var stack []string
	i := 0
	n := len(code)
	rootClosed := false

	for i < n {
		// Once the root tag is closed, nothing else is allowed
		if rootClosed {
			return false
		}

		// Check for CDATA
		if i+8 < n && code[i:i+9] == "<![CDATA[" {
			// CDATA can only appear inside a tag (stack must be non-empty)
			if len(stack) == 0 {
				return false
			}
			j := i + 9
			// Find ]]>
			endIdx := strings.Index(code[j:], "]]>")
			if endIdx == -1 {
				return false
			}
			i = j + endIdx + 3
			continue
		}

		if code[i] == '<' {
			// Check for end tag
			if i+1 < n && code[i+1] == '/' {
				// Find the closing '>'
				j := strings.IndexByte(code[i+2:], '>')
				if j == -1 {
					return false
				}
				tagName := code[i+2 : i+2+j]
				if !isValidTagName(tagName) {
					return false
				}
				// Must match the top of stack
				if len(stack) == 0 || stack[len(stack)-1] != tagName {
					return false
				}
				stack = stack[:len(stack)-1]
				i = i + 2 + j + 1
				// If stack is empty now, root tag was closed
				if len(stack) == 0 {
					rootClosed = true
				}
			} else if i+1 < n && code[i+1] == '!' {
				// Only CDATA is allowed, handled above
				return false
			} else {
				// Start tag
				j := strings.IndexByte(code[i+1:], '>')
				if j == -1 {
					return false
				}
				tagName := code[i+1 : i+1+j]
				if !isValidTagName(tagName) {
					return false
				}
				// Tag name must be between 1 and 9 uppercase letters
				if len(tagName) < 1 || len(tagName) > 9 {
					return false
				}
				stack = append(stack, tagName)
				i = i + 1 + j + 1
			}
		} else {
			i++
		}
	}

	// All tags must be closed, and root tag must have been closed exactly once
	return len(stack) == 0 && rootClosed
}

func isValidTagName(name string) bool {
	if len(name) < 1 || len(name) > 9 {
		return false
	}
	for _, ch := range name {
		if !unicode.IsUpper(ch) {
			return false
		}
	}
	return true
}
```
