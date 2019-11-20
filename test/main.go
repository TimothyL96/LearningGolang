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
// 0 < arrSize, 0 < maxSum, 0 < arr[i]
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

	calcExponential(arr, arrSize, maxSum, 0, 0, &maxPSum)

	return maxPSum
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

// AVL tree or self balancing BST
type BST struct {
}

// insert into the balanced tree
func (b *BST) insert(value int) {

}

// search the BST with log n + log n
func (b *BST) searchValueOrNextSmallest(value int) int {
	return 0
}

// self balancing tree rotate left
func (b *BST) rotateLeft() {

}

// self balancing tree rotate right
func (b *BST) rotateRight() {

}
