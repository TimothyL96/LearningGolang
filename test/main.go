package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := run(100, 5, "50 10 20 30 40")
	fmt.Println(x)
}

// Read input
func run(maxSum, arrSize int, str string) int {
	var arrStr []string
	var arr []int
	var maxPSum int

	arrStr = strings.Split(str, " ")
	for _, s := range arrStr {
		tint, _ := strconv.Atoi(s)
		arr = append(arr, tint)
	}

	fmt.Println("Max sum:", maxSum)
	fmt.Println("Input size:", arrSize)
	fmt.Println("Input:", arr)

	// calcExponential(arr, arrSize, maxSum, 0, 0, &maxPSum)
	calcNLogN(arr, arrSize, maxSum, &maxPSum)

	return maxPSum
}

// AVL tree or self balancing BST
type BST struct {
}

func calcNLogN(arr []int, arrSize, maxSum int, maxPSum *int) {
	// Set values for first element
	exclude, include, prevExclude, tmpInclude := 0, 0, 0, 0

	// Set value of include using element 0
	if arr[0] <= maxSum {
		include = arr[0]
	}

	// Test:
	// include exclude  // BST
	// 50 0     // empty    // i = 0, v = 50
	// 10,50    // 50, 0    // i = 1, v = 10
	// 70,50    // 50, 0, 10,   // i = 2, v = 20
	// 80,70    // 50, 0, 10, 70, 20    // i = 3, v = 30
	// *,80
	// tmpInclude = 110 > complementOfInclude = 60 > nextSmallestValue = 50 > tmpInclude = max(40,40+50) > 90,80    // 50, 0, 10, 70, 20, 80, 30     // i = 4, v = 40

	// Loop through array once, starting with 2nd element, i = 1 *** t = n
	for _, i := range arr[1:] {
		prevExclude = exclude
		exclude = max(exclude, include)
		tmpInclude = max(i, prevExclude+i)
		0
		if tmpInclude > maxSum {
			// if i > maxSum, set tmpInclude to 0

			complementOfInclude := maxSum - tmpInclude

			// Search for the value or the next smallest value of complementOfInclude *** t = log n
			// for next smallest value, if no left child, compare node value and parent node value and return parent value if smaller than current node value
			// else return 0
			// if there's left child, get max value in the tree with the left child as root node, *** t = h where h = height of the BST and O(h) = log n
			// go 1 left child down and right child down all the way

			// If no elements exists, tmpInclude = current value, i and i must be <= maxSum
			// If exist, tmpInclude = max(i, value in BST+i)
		}

		// Add previous include to the BST *** t = log n > Including outer loop = n * log n
		// Add previous exclude to the BST *** t = log n
		// Add previous value, i-1 to the BST *** t = log n

		// Set tmpInclude to current include
	}
	// Total t = n * (log n + log n + log n + log n + log n + log n) = n * (5 * log n)
	// t = O(n * log n)

	*maxPSum = max(exclude, include)
}

// algorithm
func calcExponential(arr []int, arrSize, maxSum, curIndex, curValue int, maxPSum *int) {
	for i := curIndex; i < arrSize; i++ {
		newCurValue := curValue + arr[i]
		if *maxPSum < newCurValue && newCurValue <= maxSum {
			*maxPSum = newCurValue
		}
		if *maxPSum == maxSum {
			return
		}

		calcExponential(arr, arrSize, maxSum, i+2, newCurValue, maxPSum)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
