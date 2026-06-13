class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix or not matrix[0]:
            return False
        for i in range(len(matrix)):
            left, right = 0, len(matrix[i]) - 1
            while left <= right:
                middle = left + (right - left) // 2
                if matrix[i][middle] == target:
                    return True
                elif matrix[i][middle] < target:
                    left = middle + 1
                else:
                    right = middle - 1
        return False