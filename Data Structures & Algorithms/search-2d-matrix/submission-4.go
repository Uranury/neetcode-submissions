func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	for i < len(matrix) {
		if target > matrix[i][len(matrix[i])-1] {
			i++
		} else {
			start, end := 0, len(matrix[i]) - 1
			for start <= end {
				middle := (start + end) / 2
				if matrix[i][middle] == target {
					return true
				} else if matrix[i][middle] < target {
					start = middle + 1
				} else {
					end = middle - 1 
				}
			}
			break
		}
	}
	return false
}
