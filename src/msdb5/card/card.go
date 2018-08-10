package card

import "strconv"

// Card type
type Card struct {
	number uint8
	seed   Seed
}

// ByID func
func ByID(id int) *Card {
	a := uint8(id % 10)
	b := Seed(id / 10)
	return &Card{number: a, seed: b}
}

func (card *Card) points() uint8 {
	switch card.number {
	case 1:
		return 11
	case 3:
		return 10
	case 8:
		return 2
	case 9:
		return 3
	case 10:
		return 4
	default:
		return 0
	}
}

func (card Card) String() string {
	return "(" + strconv.Itoa(int(card.number)) + " of " + card.seed.String() + ")"
}