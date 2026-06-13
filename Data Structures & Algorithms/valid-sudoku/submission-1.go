func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		seenRow := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := seenRow[board[i][j]]; ok {
				return false
			}
			seenRow[board[i][j]] = struct{}{}
		}

		seenCol := make(map[byte]struct{})
		for k := 0; k < 9; k++ {
			if board[k][i] == '.' {
				continue
			}
			if _, ok := seenCol[board[k][i]]; ok {
				return false
			}
			seenCol[board[k][i]] = struct{}{}
		}

		seenSquare := make(map[byte]struct{})
		for l := 0; l < 9; l++ {
			row := 3*(i/3) + l/3
			col := 3*(i%3) + l%3
			if board[row][col] == '.' {
				continue
			}

			if _, ok := seenSquare[board[row][col]]; ok {
				return false
			}

			seenSquare[board[row][col]] = struct{}{}
		}
	}

	return true
}
