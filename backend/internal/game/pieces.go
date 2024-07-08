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

func (p *Piece) ValidMoves(board *Board, start Position) []Position {
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
	return nil
}
