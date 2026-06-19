package main

// LeetCode #2726: Calculator with Method Chaining
// https://leetcode.com/problems/calculator-with-method-chaining/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Calculator with method chaining.

import "fmt"

func main() {
	result := CalculatorWithMethodChaining(10).add(5).subtract(7).getResult()
	fmt.Println(result)
}

type Calculator struct {
	value int
}

func (c *Calculator) add(val int) *Calculator {
	c.value += val
	return c
}

func (c *Calculator) subtract(val int) *Calculator {
	c.value -= val
	return c
}

func (c *Calculator) multiply(val int) *Calculator {
	c.value *= val
	return c
}

func (c *Calculator) divide(val int) *Calculator {
	c.value /= val
	return c
}

func (c *Calculator) power(val int) *Calculator {
	result := 1
	for i := 0; i < val; i++ {
		result *= c.value
	}
	c.value = result
	return c
}

func (c *Calculator) getResult() int {
	return c.value
}

func CalculatorWithMethodChaining(initialValue int) *Calculator {
	return &Calculator{value: initialValue}
}
