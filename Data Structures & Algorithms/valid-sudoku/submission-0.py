class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for row in board:
            rowSeen = set()
            for cell in row:
                if cell in rowSeen and cell != ".":
                    return False
                rowSeen.add(cell) 
        for i in range(9):
            colSeen = set()
            for row in board:
                if row[i] in colSeen and row[i] != ".":
                    return False
                colSeen.add(row[i])
        for row in range(0, 9, 3):
            for col in range(0, 9, 3):
                subgrid = set()
                for r in range(row, row + 3):
                    for c in range(col, col + 3):
                        if board[r][c] in subgrid and board[r][c] != ".":
                            return False
                        subgrid.add(board[r][c])
        return True