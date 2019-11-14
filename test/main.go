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

func calcNLogN(arr []int, arrSize, maxSum int, maxPSum *int) {
	// AVL tree or self balancing BST
	type BST struct {
	}

	// Set values for first element
	exclude := 0
	include := 0
	prevExclude := 0
	tmpInclude := 0
	if arr[0] <= maxSum {
		include = arr[0]
	}

	// Loop through array once, starting with 2nd element, i = 1
	// First layer -  n
	for _, i := range arr[1:] {
		prevExclude = exclude
		exclude = max(exclude, exclude)
		tmpInclude = max(i, prevExclude+i)

		if tmpInclude > maxSum {
			complementOfInclude := maxSum - tmpInclude

			// Search for the value or the next smallest value of complementOfInclude
			// If no elements exists, tmpInclude = current value, i
			// If exist, tmpInclude = max(i, value in BST+i)
		}

		// Add previous include to the BST
		// Add current value, i to the BST
		// Add previous exclude to the BST
		// Add previous index to the BST

		// Set tmpInclude to current include
	}

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
