package game

type Game struct {
	Board *Board
	Turn  Color
}

func NewGame() *Game {

	// Initialize a new game board and the first player turn
	game := &Game{
		Board: NewBoard(),
		Turn:  P1,
	}

	game.SetupPieces()
	return game

}

// function to setup the initial gmae board with the pieces
// following a 3d coordinate system
func (g *Game) SetupPieces() {

	//
	//												//
	// Set up the pieces for the first player first //
	//												//
	//

	// Set the pawns up first
	g.Board.SetPiece(Position{Q: 0, R: 1, S: -1}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: -1, R: 2, S: -1}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: 1, R: 1, S: -2}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: -2, R: 3, S: -1}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: 2, R: 1, S: -3}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: -3, R: 4, S: -1}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: 3, R: 1, S: -4}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: -4, R: 5, S: -1}, &Piece{T: Pawn, C: P1})
	g.Board.SetPiece(Position{Q: 4, R: 1, S: -5}, &Piece{T: Pawn, C: P1})

	// Set up the bishops
	for i := 3; i < 6; i++ {
		g.Board.SetPiece(Position{Q: 0, R: -i, S: i}, &Piece{T: Bishop, C: P1})
	}

	// Set up the rooks
	g.Board.SetPiece(Position{Q: -3, R: 5, S: -2}, &Piece{T: Rook, C: P1})
	g.Board.SetPiece(Position{Q: 3, R: 2, S: -5}, &Piece{T: Rook, C: P1})

	// Set up the Knight
	g.Board.SetPiece(Position{Q: -2, R: 5, S: -3}, &Piece{T: Knight, C: P1})
	g.Board.SetPiece(Position{Q: 2, R: 3, S: -5}, &Piece{T: Knight, C: P1})

	// Set up the Queen
	g.Board.SetPiece(Position{Q: -1, R: 5, S: -4}, &Piece{T: Queen, C: P1})

	// Set up the King
	g.Board.SetPiece(Position{Q: 1, R: 4, S: -5}, &Piece{T: King, C: P1})

	//
	//											//
	// Set up the Pieces for the Second Player	//
	//											//
	//

	// Set the pawns up first
	g.Board.SetPiece(Position{Q: 0, R: -1, S: 1}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: 1, R: -2, S: 1}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: -1, R: -1, S: 2}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: 2, R: -3, S: 1}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: -2, R: -1, S: 3}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: 3, R: -4, S: 1}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: -3, R: -1, S: 4}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: 4, R: -5, S: 1}, &Piece{T: Pawn, C: P2})
	g.Board.SetPiece(Position{Q: -4, R: -1, S: 5}, &Piece{T: Pawn, C: P2})

	// Set up the bishops
	for i := -5; i <= -3; i++ {
		g.Board.SetPiece(Position{Q: 0, R: -1 * i, S: i}, &Piece{T: Bishop, C: P2})
	}

	// Set up the rooks
	g.Board.SetPiece(Position{Q: 3, R: -5, S: 2}, &Piece{T: Rook, C: P2})
	g.Board.SetPiece(Position{Q: -3, R: -2, S: 5}, &Piece{T: Rook, C: P2})

	// Set up the Knight
	g.Board.SetPiece(Position{Q: 2, R: -5, S: 3}, &Piece{T: Knight, C: P2})
	g.Board.SetPiece(Position{Q: -2, R: -3, S: 5}, &Piece{T: Knight, C: P2})

	// Set up the King
	g.Board.SetPiece(Position{Q: 1, R: -5, S: 4}, &Piece{T: King, C: P2})

	// Set up the Queen
	g.Board.SetPiece(Position{Q: -1, R: -4, S: 5}, &Piece{T: Queen, C: P2})

}
