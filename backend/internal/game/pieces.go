package game

// This helps store the type of a string in an identifable manner
type PieceType int

const (
	King   PieceType = 0
	Queen  PieceType = 1
	Rook   PieceType = 2
	Bishop PieceType = 3
	Knight PieceType = 4
	Pawn   PieceType = 5
)

// This helps keep the color adjustable but the player order fixed
type Color int

const (
	P1 Color = 1
	P2 Color = 2
)

type Piece struct {
	T PieceType
	C Color
}

type Position struct {
	Q int
	R int
	S int
}

type Moves struct {
	positions []Position
	captures  []Position
}

func (p *Piece) ValidMoves(board *Board, start Position) Moves {
	switch p.T {
	case King:
		return validKingMoves(board, start)
	case Queen:
		return validQueenMoves(board, start)
	case Rook:
		return validRookMoves(board, start)
	case Bishop:
		return validBishopMoves(board, start)
	case Knight:
		return validKnightMoves(board, start)
	case Pawn:
		return validPawnMoves(board, start, p.C)
	}
	return Moves{positions: nil, captures: nil}
}

func validPawnMoves(board *Board, start Position, c Color) Moves {
	var moves Moves
	if c == P1 {

		// Check for a single move forward
		f1 := Position{Q: start.Q, R: start.R - 1, S: start.S + 1}
		if board.IsValidPosition(f1) && board.GetPiece(f1) == nil {
			moves.positions = append(moves.positions, f1)

			// The only situation where the pawn can move 2 spaces is when it can move at least once
			f2 := Position{Q: start.Q, R: start.R - 2, S: start.S + 2}
			if board.IsValidPosition(f2) && board.GetPiece(f2) == nil {
				moves.positions = append(moves.positions, f2)
			}
		}

		// Check capture points
		c1 := Position{Q: start.Q - 1, R: start.R, S: start.S + 1}
		if board.IsValidPosition(c1) && board.GetPiece(c1) != nil {
			moves.captures = append(moves.captures, c1)
		}

		c2 := Position{Q: start.Q + 1, R: start.R - 1, S: start.S}
		if board.IsValidPosition(c2) && board.GetPiece(c2) != nil {
			moves.captures = append(moves.captures, c2)
		}

	} else {

		f1 := Position{Q: start.Q, R: start.R + 1, S: start.S - 1}
		if board.IsValidPosition(f1) && board.GetPiece(f1) == nil {
			moves.positions = append(moves.positions, f1)

			f2 := Position{Q: start.Q, R: start.R + 2, S: start.S - 2}
			if board.IsValidPosition(f2) && board.GetPiece(f2) == nil {
				moves.positions = append(moves.positions, f2)
			}
		}

		c1 := Position{Q: start.Q - 1, R: start.R + 1, S: start.S}
		if board.IsValidPosition(c1) && board.GetPiece(c1) != nil {
			moves.captures = append(moves.captures, c1)
		}

		c2 := Position{Q: start.Q + 1, R: start.R, S: start.S - 1}
		if board.IsValidPosition(c2) && board.GetPiece(c2) != nil {
			moves.captures = append(moves.captures, c2)
		}
	}

	return moves
}
