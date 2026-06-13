class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix or not matrix[0]:
            return False
        rows, cols = 0, len(matrix) - 1
        ROW = 0
        while rows <= cols:
            middle = rows + (cols - rows) // 2
            if matrix[middle][0] <= target <= matrix[middle][-1]:
                ROW = middle
                break
            elif target > matrix[middle][-1]:
                rows = middle + 1
            else:
                cols = middle - 1
        left, right = 0, len(matrix[ROW]) - 1
        while left <= right:
            middle = left + (right - left) // 2
            if matrix[ROW][middle] == target:
                return True
            elif matrix[ROW][middle] > target:
                right = middle - 1
            else:
                left = middle + 1
        return False