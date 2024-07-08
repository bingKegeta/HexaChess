package game

func IsCheck(board *Board, kingColor Color) bool {
	var kingPos Position
	for pos, piece := range board.Cells {
		if piece != nil && piece.T == King && piece.C == kingColor {
			kingPos = pos
			break
		}
	}

	for pos, piece := range board.Cells {
		if piece != nil && piece.C != kingColor {
			moves := piece.ValidMoves(board, pos)
			for _, move := range moves {
				if move == kingPos {
					return true
				}
			}
		}
	}
	return false
}

func IsCheckmate(board *Board, kingColor Color) bool {
	if !IsCheck(board, kingColor) {
		return false
	}

	for pos, piece := range board.Cells {
		if piece != nil && piece.C == kingColor {
			moves := piece.ValidMoves(board, pos)
			for _, move := range moves {
				tempBoard := CopyBoard(board)
				tempBoard.MovePiece(pos, move)
				if !IsCheck(tempBoard, kingColor) {
					return false
				}
			}
		}
	}
	return true
}

func CopyBoard(board *Board) *Board {
	newBoard := NewBoard()
	for pos, piece := range board.Cells {
		newBoard.Cells[pos] = piece
	}
	return newBoard
}
