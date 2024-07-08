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
	return pos.Q+pos.R+pos.S == 0 // In a cube coordinate system, q + r + s should always be 0
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

func (b *Board) SetPiece(pos Position, piece *Piece) {
	if b.IsValidPosition(pos) {
		b.Cells[pos] = piece
	}
}
