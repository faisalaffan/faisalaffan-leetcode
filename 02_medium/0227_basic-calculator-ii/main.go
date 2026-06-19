package main

// LeetCode #227: Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func calculate(s string) int {
	stack := []int{}
	num := 0
	op := '+'

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if s[i] == ' ' && i != len(s)-1 {
				continue
			}
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = ch
			num = 0
		}
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}
