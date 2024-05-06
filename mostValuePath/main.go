package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
	prev  **BinaryNode
}

func findMax(root *BinaryNode) int {
	if root == nil {
		return 0
	}
	max := 0

	root.preOrderTraverse([]int{}, &max)

	return max
}

func (root *BinaryNode) preOrderTraverse(currentpath []int, max *int) {
	if root == nil {
		return
	}

	currentpath = append(currentpath, root.value)

	if root.left == nil && root.right == nil {
		sum := sumNumber(currentpath)
		if sum > *max {
			*max = sum
		}
		sum = 0
	} else {
		root.left.preOrderTraverse(currentpath, max)
		root.right.preOrderTraverse(currentpath, max)
	}

	currentpath = currentpath[:len(currentpath)-1]
}

func sumNumber(data []int) int {
	sum := 0
	for index := range len(data) {
		sum += data[index]
	}
	return sum
}

func readJSONfile() ([][]int, error) {
	file, err := os.Open("./tree1.json")
	if err != nil {
		return nil, fmt.Errorf("Error Opening File: %w", err)
	}

	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error IO Read File: %w", err)
	}

	var data [][]int
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshal JSON: %w", err)
	}

	return data, nil
}

func arrayToPymarid(array [][]int) *BinaryNode {
	if len(array) == 0 {
		return nil
	}

	root := &BinaryNode{value: array[0][0]}
	var rootPointer **BinaryNode
	rootPointer = &root
	var sharedPointer **BinaryNode
	var mostLeftNode *BinaryNode

	for level := 0; level < len(array)-1; level++ {
		if level == 0 {
			root.left = &BinaryNode{value: array[level+1][0]}

			root.left.prev = &root
			root.right = &BinaryNode{value: array[level+1][1]}

			mostLeftNode = root.left
			rootPointer = &mostLeftNode
			continue
		}

		for index := range len(array[level]) {
			// most left
			if index == 0 {
				tempRoot := *rootPointer
				tempRoot.left = &BinaryNode{value: array[level+1][index]}
				tempRoot.left.prev = &tempRoot // for walking back
				mostLeftNode = tempRoot.left

				// reserve shared node
				sharedNode := &BinaryNode{value: array[level+1][index+1]}
				sharedPointer = &sharedNode
				tempRoot.right = *sharedPointer

				// point to right adjacent node
				if ok := tempRoot.prev; ok != nil {
					previousNode := *tempRoot.prev
					rightAdjacentNode := previousNode.right
					rootPointer = &rightAdjacentNode
				}
			} else {
				tempRoot := *rootPointer
				tempRoot.left = *sharedPointer
				tempRoot.left.prev = &tempRoot

				sharedNode := &BinaryNode{value: array[level+1][index+1]}
				sharedPointer = &sharedNode
				tempRoot.right = *sharedPointer

				if ok := tempRoot.prev; ok != nil {
					previousNode := *tempRoot.prev
					rightAdjacentNode := previousNode.right
					rootPointer = &rightAdjacentNode
				}
			}
		}

		if mostLeftNode != nil {
			rootPointer = &mostLeftNode
		}
	}
	return root
}

func main() {
	actual, _ := readJSONfile()
	pym := arrayToPymarid(actual)
	max := findMax(pym)
	fmt.Println("this is max value among all paths :", max)
}
