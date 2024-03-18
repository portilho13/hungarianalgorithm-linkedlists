package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data     int
	row      int
	column   int
	isMarked bool
	next     *Node
}

func newNode(data, row, column int) *Node {
	return &Node{data: data, row: row, column: column}
}

func insertNode(head **Node, data, row, column int) {
	newNode := newNode(data, row, column)
	if newNode == nil {
		fmt.Println("Error creating new node")
		return
	}
	if *head == nil {
		*head = newNode
	} else {
		temp := *head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

func createCopyNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := newNode(head.data, head.row, head.column)
	current := head.next
	newCurrent := newHead

	for current != nil {
		newNode := newNode(current.data, current.row, current.column)
		newCurrent.next = newNode
		newCurrent = newCurrent.next
		current = current.next
	}
	return newHead
}

func printList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("Data: %d Row: %d Column: %d isMarked: %t\n", temp.data, temp.row, temp.column, temp.isMarked)
		temp = temp.next
	}
}

func readFile(head **Node) (int, int) {
	file, err := os.Open("ficheiro.txt")
	if err != nil {
		fmt.Println("File reading error:", err)
		return 0, 0
	}
	defer file.Close()
	var row, col int
	var numRows, numCols int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStrs := strings.Split(scanner.Text(), ";")
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			insertNode(head, num, row, col)
			col++
		}
		row++
		numCols = col
		col = 0
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}
	numRows = row
	return numRows, numCols
}

func getMaximumValue(head *Node) int {
	temp := head
	max := temp.data
	for temp != nil {
		if temp.data > max {
			max = temp.data
		}
		temp = temp.next
	}
	return max
}

func applyMaximum(head *Node) {
	temp := head
	max := getMaximumValue(head)
	for temp != nil {
		temp.data = max - temp.data
		temp = temp.next
	}
}

func getMinimumInRow(head *Node, row int) int {
	temp := head
	min := temp.data
	for temp != nil {
		if temp.row == row && temp.data < min {
			min = temp.data
		}
		temp = temp.next
	}
	return min
}

func getMinimumInColumn(head *Node, col int) int {
	temp := head
	min := temp.data
	for temp != nil {
		if temp.column == col && temp.data < min {
			min = temp.data
		}
		temp = temp.next
	}
	return min
}

func findZerosInRow(head *Node, row int) int {
	temp := head
	count := 0
	for temp != nil {
		if temp.row == row && temp.data == 0 {
			count++
		}
		temp = temp.next
	}
	return count
}

func findZerosInColumn(head *Node, column int) int {
	temp := head
	count := 0
	for temp != nil {
		if temp.column == column && temp.data == 0 {
			count++
		}
		temp = temp.next
	}
	return count
}

func selectZeros(head *Node, maxZerosRow int, maxZerosColumn int, rowMarked []bool, columnMarked []bool) {
	var temp *Node
	if maxZerosColumn > maxZerosRow {
		temp = head
		for temp != nil {
			numberOfZerosInColumn := findZerosInColumn(temp, temp.column)
			if temp.data == 0 && numberOfZerosInColumn > 1 && !columnMarked[temp.column] && !rowMarked[temp.row] {
				temp.isMarked = true
				columnMarked[temp.column] = true
			}
			temp = temp.next
		}

		temp = head
		for temp != nil {
			numberOfZerosInRow := findZerosInRow(temp, temp.row)
			if temp.data == 0 && numberOfZerosInRow > 1 && !rowMarked[temp.row] && !temp.isMarked {
				temp.isMarked = true
				rowMarked[temp.row] = true
			}
			temp = temp.next
		}
	} else {
		temp = head
		for temp != nil {
			numberOfZerosInRow := findZerosInRow(temp, temp.row)
			if temp.data == 0 && numberOfZerosInRow > 1 && !rowMarked[temp.row] && !temp.isMarked {
				temp.isMarked = true
				rowMarked[temp.row] = true
			}
			temp = temp.next
		}

		temp = head
		for temp != nil {
			numberOfZerosInColumn := findZerosInColumn(temp, temp.column)
			if temp.data == 0 && numberOfZerosInColumn > 1 && !columnMarked[temp.column] && !rowMarked[temp.row] {
				temp.isMarked = true
				columnMarked[temp.column] = true
			}
			temp = temp.next
		}
	}
}

func findPermutations(head *Node, matrix [][]int, rows, cols int, assigned [][]int, currentRow, selectedZeros int) int {
    count := 0
    maxSum := 0

    // Base case: If all zeros are selected
    if selectedZeros == rows {
        // Check if there's a unique 0 per row and column
        unique := true
        for i := 0; i < rows; i++ {
            rowCount, colCount := 0, 0
            for j := 0; j < cols; j++ {
                rowCount += assigned[i][j] // Count the number of zeros in the row
                colCount += assigned[j][i] // Count the number of zeros in the column
            }
            if rowCount != 1 || colCount != 1 {
                unique = false
                break
            }
        }

        // If the permutation is unique, calculate the sum
        if unique {
            temp := head
            for i := 0; i < rows; i++ {
                for j := 0; j < cols; j++ {
                    if assigned[i][j] == 1 {
                        for temp != nil {
                            if temp.row == i && temp.column == j {
                                count += temp.data // Add the value to the sum
                                break
                            }
                            temp = temp.next // Move to the next node
                        }
                        temp = head // Reset temp to the head for the next iteration
                    }
                }
            }

            // Update the maximum sum if the current sum is greater
            if count > maxSum {
                maxSum = count
            }
        }
    }

    // Recursive case: Explore all possible assignments
    if currentRow < rows {
        for col := 0; col < cols; col++ {
            if matrix[currentRow][col] == 0 && assigned[currentRow][col] == 0 { // If the current element is zero and not assigned
                assigned[currentRow][col] = 1 // Assign the zero
                // Recur for the next row and increase the count of selected zeros
                tempCount := findPermutations(head, matrix, rows, cols, assigned, currentRow, selectedZeros+1)
                // Update count and maxSum
                if tempCount > count { // If the current count is greater than the previous count
                    count = tempCount
                    if count > maxSum { // If the current count is greater than the maximum sum encountered
                        maxSum = count
                    }
                }
                assigned[currentRow][col] = 0 // Backtrack: Unassign the zero
            }
        }
        // Recur for the next row
        tempCount := findPermutations(head, matrix, rows, cols, assigned, currentRow+1, selectedZeros)
        // Update count and maxSum
        if tempCount > count {
            count = tempCount
            if count > maxSum {
                maxSum = count
            }
        }
    }
    return maxSum // Return the maximum sum encountered
}

func HungarianAlgorithm(head *Node, rows int, cols int) (int) {
	var copy, Step1, Step2 *Node
	copy = createCopyNode(head)

	tempCopy := createCopyNode(head)

	applyMaximum(copy)

	// Step1
	Step1 = createCopyNode(copy)
	temp := Step1
	for temp != nil {
		temp.data = temp.data - getMinimumInRow(copy, temp.row)
		temp = temp.next
	}

	// Step2
	Step2 = createCopyNode(Step1)
	temp = Step2
	for temp != nil {
		temp.data = temp.data - getMinimumInColumn(Step1, temp.column)
		temp = temp.next
	}

	var rowMarked = make([]bool, rows)
	var colMarked = make([]bool, cols)

	maxZeroInRow := 0
	temp = Step2
	for temp != nil {
		numberOfZerosInRow := findZerosInRow(temp, temp.row)
		if numberOfZerosInRow > maxZeroInRow {
			maxZeroInRow = numberOfZerosInRow
		}
		temp = temp.next
	}

	maxZeroInColumn := 0
	temp = Step2
	for temp != nil {
		numberOfZerosInColumn := findZerosInColumn(temp, temp.column)
		if numberOfZerosInColumn > maxZeroInColumn {
			maxZeroInColumn = numberOfZerosInColumn
		}
		temp = temp.next
	}

	selectZeros(Step2, maxZeroInRow, maxZeroInColumn, rowMarked, colMarked)
	temp = Step2
	for temp != nil {
		if temp.data == 0 && !rowMarked[temp.row] && !colMarked[temp.column] {
			rowMarked[temp.row] = true
		}
		temp = temp.next
	}

	numberOfCrossedRows := 0
	for i := 0; i < rows; i++ {
		if (rowMarked[i]) {
			numberOfCrossedRows++
		}
	}

	numberOfCrossedColumns := 0
	for i := 0; i < cols; i++ {
		if (colMarked[i]) {
			numberOfCrossedColumns++
		}
	}

	sum := numberOfCrossedRows + numberOfCrossedColumns

	if (sum != rows) {
		for sum != rows {
			temp = Step2
			min := temp.data
			for temp != nil {
				if (!rowMarked[temp.row] && !colMarked[temp.column]) {
					if (temp.data < min) {
						min = temp.data
					}
				}
				temp = temp.next
			}

			temp = Step2

			for temp != nil {
				if (rowMarked[temp.row] && colMarked[temp.column]) {
					temp.data += min
				} else if !rowMarked[temp.row] && !colMarked[temp.column] {
					temp.data -= min
				}
				temp = temp.next
			}

			for i := 0; i < rows; i++ {
				rowMarked[i] = false
				colMarked[i] = false
			}

			temp = Step2

			for temp != nil {
				temp.isMarked = false
				temp = temp.next
			}

			temp = Step2
			maxZeroInRow = 0

			for temp != nil {
				numberOfZerosInRow := findZerosInRow(temp, temp.row)
				if (numberOfZerosInRow > maxZeroInRow) {
					maxZeroInRow = numberOfZerosInRow
				}
				temp = temp.next
			}

			temp = Step2
			maxZeroInColumn = 0
			
			for temp != nil {
				numberOfZerosInColumn := findZerosInColumn(temp, temp.column)

				if numberOfZerosInColumn > maxZeroInColumn {
					maxZeroInColumn = numberOfZerosInColumn
				}
				temp = temp.next
			}

			selectZeros(Step2, maxZeroInRow, maxZeroInColumn, rowMarked, colMarked)

			temp = Step2

			for temp != nil {
				if temp.data == 0 && !rowMarked[temp.row] && !colMarked[temp.column] {
					rowMarked[temp.row] = true
				}
				temp = temp.next
			}

			numberOfCrossedRows := 0
			for i := 0; i < rows; i++ {
				if (rowMarked[i]) {
					numberOfCrossedRows++
				}
			}
		
			numberOfCrossedColumns := 0
			for i := 0; i < cols; i++ {
				if (colMarked[i]) {
					numberOfCrossedColumns++
				}
			}
			sum = numberOfCrossedRows + numberOfCrossedColumns
		}
	}

	printList(Step2)
	for i := 0; i < rows; i++ {
		rowMarked[i] = false
		colMarked[i] = false
	}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	temp = Step2
	for temp != nil {
		matrix[temp.row][temp.column] = temp.data
		temp = temp.next
	}

	assigned := make([][]int, rows)
	for i := range assigned {
		assigned[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			assigned[i][j] = 0
		}
	}
	count := findPermutations(tempCopy, matrix, rows, cols, assigned, 0, 0)
	return count

}

func main() {
	var head *Node
	rows, cols := readFile(&head)	
	count := HungarianAlgorithm(head, rows, cols)
	fmt.Println("Maximum sum:", count)
}
