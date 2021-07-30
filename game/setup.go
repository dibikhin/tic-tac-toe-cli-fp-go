package game

import (
	. "tictactoe/internal"
)

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...reader) (Game, error) {
	alt, err := extractReader(rs)
	if err != nil {
		return DeadGame(), err
	}
	gam, err := makeGame(DefaultReader, alt)
	if err != nil {
		return DeadGame(), err
	}
	PrintLogo(gam.Logo)

	defer gam.Print()
	p1, p2, err := gam.ChooseMarks()
	if err != nil {
		return DeadGame(), err
	}
	return SetPlayers(gam, p1, p2), nil
}

// Private

// Factory, Pure
func makeGame(def, alt reader) (Game, error) {
	gam := NewGame()
	switch {
	case alt != nil:
		return SetReader(gam, alt)
	default:
		return SetReader(gam, def)
	}
}
