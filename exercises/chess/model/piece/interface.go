// Package piece will handle game pieces.
package piece

import (
	"fmt"

	"github.com/Horsefactor/4eail40_2020_2021/exercises/chess/model/coord"
	"github.com/Horsefactor/4eail40_2020_2021/exercises/chess/model/player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns the appartenance.
	Color() player.Color
	// Moves returns a set of valid move.
	Moves(isCapture bool) map[coord.ChessCoordinates]bool
}
