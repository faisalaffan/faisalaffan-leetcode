package main

import (
	"fmt"
	"strings"
)

// LeetCode #1597: Build Binary Expression Tree From Infix Expression
// https://leetcode.com/problems/build-binary-expression-tree-from-infix-expression/
// Difficulty: Hard [Paid]
//
// Build a binary expression tree from an infix expression string.
// The tree follows standard operator precedence: * and / bind tighter than + and -.
// The tree nodes are:
//
//	type Node struct {
//		Val   byte   // '0'-'9' for digits, '+', '-', '*', '/' for operators
//		Left  *Node
//		Right *Node
//	}
//
// Operators are stored in internal nodes; operands (digits) are leaves.

// Node represents a node in the binary expression tree.
type Node struct {
	Val   byte
	Left  *Node
	Right *Node
}

// expTree builds a binary expression tree from the infix expression s.
func expTree(s string) *Node {
	// Remove spaces
	s = strings.ReplaceAll(s, " ", "")

	// Shunting-yard: convert infix to postfix (RPN)
	var postfix []byte
	var ops []byte

	prec := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			postfix = append(postfix, ch)
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1] // pop '('
		} else {
			// operator
			for len(ops) > 0 && ops[len(ops)-1] != '(' &&
				prec[ops[len(ops)-1]] >= prec[ch] {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, ch)
		}
	}

	for len(ops) > 0 {
		postfix = append(postfix, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	// Build expression tree from postfix using a stack
	var stack []*Node
	for _, token := range postfix {
		if token >= '0' && token <= '9' {
			stack = append(stack, &Node{Val: token})
		} else {
			// Operator: pop two operands
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Node{
				Val:   token,
				Left:  left,
				Right: right,
			})
		}
	}

	if len(stack) != 1 {
		return nil
	}
	return stack[0]
}

// inorder returns the infix representation (with parentheses to verify structure).
func inorder(root *Node) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return string(root.Val)
	}
	return "(" + inorder(root.Left) + string(root.Val) + inorder(root.Right) + ")"
}

// postorder returns the postfix representation.
func postorder(root *Node) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return string(root.Val)
	}
	return postorder(root.Left) + postorder(root.Right) + string(root.Val)
}

func main() {
	// Example 1:
	// Input: "3*4-2*5"
	// Output: - (subtree: * and *)
	// Tree:
	//        -
	//      /   \
	//     *     *
	//    / \   / \
	//   3   4 2   5
	tree1 := expTree("3*4-2*5")
	fmt.Println("Infix:  ", inorder(tree1))
	fmt.Println("Postfix:", postorder(tree1))
	// Expected infix: ((3*4)-(2*5))
	// Expected postfix: 34*25*-

	// Example 2:
	// Input: "2-3/(5*2)+1"
	tree2 := expTree("2-3/(5*2)+1")
	fmt.Println("Infix:  ", inorder(tree2))
	fmt.Println("Postfix:", postorder(tree2))

	// Simple: "1+2"
	tree3 := expTree("1+2")
	fmt.Println("Infix:  ", inorder(tree3))   // (1+2)
	fmt.Println("Postfix:", postorder(tree3)) // 12+

	// Parentheses: "(1+2)*3"
	tree4 := expTree("(1+2)*3")
	fmt.Println("Infix:  ", inorder(tree4))   // ((1+2)*3)
	fmt.Println("Postfix:", postorder(tree4)) // 12+3*
}
