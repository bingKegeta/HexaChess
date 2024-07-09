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
	// Get the piece to be moved instead of passing it as a param
	piece := b.GetPiece(start)
	if piece == nil {
		//* WRONG FUNCION CALL: The piece did not exist in the intended position
		return false
	}

	if b.GetPiece(end) != nil {
		//* WRONG FUNCTION CALL: Piece already present here, should use capture method instead
		return false
	}

	// Finally set the piece to the position (assuming the coord was in bounds ofc)
	if b.SetPiece(end, piece) {
		return true
	} else {
		// In case it is out of bounds
		return false
	}
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
