package main

// LeetCode #224: Basic Calculator
// https://leetcode.com/problems/basic-calculator/
// Difficulty: Hard

import "fmt"

func calculate(s string) int {
	stack := make([]int, 0)
	result := 0
	sign := 1
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
			i--
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, result, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result = stack[len(stack)-2] + stack[len(stack)-1]*result
			stack = stack[:len(stack)-2]
		}
	}

	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
