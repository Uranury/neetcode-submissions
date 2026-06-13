func searchMatrix(matrix [][]int, target int) bool {
	start, end := 0, len(matrix) - 1
	row := -1 
	for start <= end {
		middle := (start + end) / 2
		if target >= matrix[middle][0] && target <= matrix[middle][len(matrix[middle])-1] {
			row = middle
			break 
		} else if target > matrix[middle][len(matrix[middle])-1] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	if row == -1 {
		return false
	}

	start, end = 0, len(matrix[row]) - 1
	for start <= end {
		middle := (start + end) / 2
		if matrix[row][middle] == target {
			return true
		} else if matrix[row][middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return false
}
