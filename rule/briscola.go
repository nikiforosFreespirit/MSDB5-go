package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {
	onlyOtherCardIsBriscola := !isBriscola(base, briscola) && isBriscola(other, briscola)
	otherCardIsBiggerAndOfSameSeed := doesOtherCardWin(base, other)
	return onlyOtherCardIsBriscola || otherCardIsBiggerAndOfSameSeed
}

func doesOtherCardWin(base, other card.ID) bool {
	areSeedDifferent := base.Seed() != other.Seed()
	isOtherGreaterOnPoints := base.Points() < other.Points()
	isOtherGreaterOnNumberOnly := base.Points() == other.Points() && base.Number() < other.Number()
	return !areSeedDifferent && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
}

func isBriscola(card card.ID, briscola card.Seed) bool {
	return card.Seed() == briscola
}