package game

type Board struct {
	Cells map[Position]*Piece
}

func NewBoard() *Board {
	return &Board{
		Cells: make(map[Position]*Piece),
	}
}

func (b *Board) IsValidPosition(pos Position) bool {
	return (pos.Q+pos.R+pos.S == 0) && (pos.Q <= 5 && pos.Q >= -5) && (pos.R <= 5 && pos.R >= -5) && (pos.S <= 5 && pos.S >= -5) // In a cube coordinate system, q + r + s should always be 0
}

func (b *Board) GetPiece(pos Position) *Piece {
	return b.Cells[pos]
}

func (b *Board) MovePiece(start, end Position) bool {
	piece := b.GetPiece(start)
	if piece == nil {
		return false
	}

	validMoves := piece.ValidMoves(b, start)
	for _, move := range validMoves {
		if move == end {
			b.Cells[end] = piece
			delete(b.Cells, start)
			return true
		}
	}
	return false
}

func (b *Board) SetPiece(pos Position, piece *Piece) bool {
	if b.IsValidPosition(pos) {
		b.Cells[pos] = piece
		return true
	}
	return false
}

// This function removes the piece in the position and sets the new piece in that position
// TODO Keep track of the captured pieces
func (b *Board) Capture(pos Position, newPiece Piece) bool {
	//* Bug if this gets passed
	if b.GetPiece(pos) == nil {
		return false
	}

	if b.SetPiece(pos, &newPiece) {
		return true
	} else {
		return false
	}
}
