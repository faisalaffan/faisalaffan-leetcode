package main

// LeetCode #478: Generate Random Point in a Circle
// https://leetcode.com/problems/generate-random-point-in-a-circle/
// Difficulty: Medium
// Time: O(1) per call
// Space: O(1)

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	sol := Constructor([]float64{1, 0, 0})
	// Generate a few random points
	for i := 0; i < 3; i++ {
		p := sol.RandPoint()
		fmt.Printf("%.4f %.4f\n", p[0], p[1])
	}
}

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radiusAndCenter []float64) Solution {
	return Solution{
		radius:  radiusAndCenter[0],
		xCenter: radiusAndCenter[1],
		yCenter: radiusAndCenter[2],
	}
}

func (s *Solution) RandPoint() []float64 {
	// Use random angle and random radius (sqrt for uniform distribution)
	angle := rand.Float64() * 2 * math.Pi
	r := s.radius * math.Sqrt(rand.Float64())
	x := s.xCenter + r*math.Cos(angle)
	y := s.yCenter + r*math.Sin(angle)
	return []float64{x, y}
}
