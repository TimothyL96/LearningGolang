package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preArr := []int{1, 2, 4, 5, 14, 18, 6, 7, 8, 9, 13, 10, 11, 12, 14}
	inArr := []int{5, 4, 18, 14, 2, 8, 7, 6, 9, 1, 13, 10, 12, 11, 14}

	parent := buildTree(preArr, inArr)
	postOrderStr := getPostOrder(parent)

	fmt.Println("Result:", postOrderStr)
}

func buildTree(preOrder, inOrder []int) *TreeNode {
	// Node to be used to store the first node
	var parent *TreeNode

	// Store the rightmost node
	var lastRight *TreeNode

	// Record the position traversed in preOrder
	var lastPreArrPosition int

	// Length of middle tree
	var midTreeLength int

	// Build a stack for the preOrder till the leftmost node
	var stack []int

	for k, i := range preOrder {
		if i == inOrder[0] {
			parent = &TreeNode{Val: i}
			lastPreArrPosition = k
			break
		} else {
			stack = append(stack, i)
		}
	}

	// If only one element exists in in-order array, return that node (parent variable)
	if len(inOrder) < 2 {
		return parent
	}

	// Traverse the inOrder to build the Left side of tree from the main root
	for k, i := range inOrder[1:] {
		stackLen := len(stack)

		// If stack is empty, exit loop as it should have completed Root and Left part of tree
		if stackLen > 0 {
			// Pop the stack and set current parent as Left child of current inOrder index node
			if i == stack[stackLen-1] {
				parent = &TreeNode{
					Val:  i,
					Left: parent,
				}
				stack = stack[:len(stack)-1]
				midTreeLength = 0
			} else if len(inOrder) > k+2 && len(preOrder) > 0 && stack[stackLen-1] == inOrder[k+2] {
				// If the next array is in the stack,
				// pass subarray from the first element that is not equal to the stack last node
				// to current index
				// Call recursive to create sub tree as a tree
				lastPreArrPosition++
				parent.Right = buildTree(preOrder[lastPreArrPosition:lastPreArrPosition+midTreeLength+1], inOrder[k+1-midTreeLength:])
				lastPreArrPosition = lastPreArrPosition + midTreeLength
			} else {
				// Record length of subarray not in the stack
				midTreeLength++
			}
		} else {
			break
		}
	}

	// The remaining of preOrder is the Right side of the tree
	lastRight = parent

	// Recursively call Right side of tree/array to build the tree and set it as parent Right
	if len(preOrder) > 0 {
		lastRight.Right = buildTree(preOrder[lastPreArrPosition+1:], inOrder[lastPreArrPosition+1:])
		lastRight = lastRight.Right
	}

	return parent
}

func getPostOrder(parent *TreeNode) (str string) {
	if parent == nil {
		return ""
	}

	if parent.Left != nil {
		str = getPostOrder(parent.Left)
	}

	if parent.Right != nil {
		str += getPostOrder(parent.Right)
	}

	str += strconv.Itoa(parent.Val) + " "

	return
}
