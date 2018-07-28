package msdb5

import "math"

func (a Card) Compare(b Card) int {
	c := a.compareOnSeed(&b)
	if c == 0 {
		c = a.compareOnPoints(&b)
		if c == 0 {
			c = a.compareOnNumber(&b)
		}
	}
	return c
}

func (a *Card) compareOnSeed(b *Card) int {
	seedForA := float64(a.seed)
	seedForB := float64(b.seed)
	return int(math.Abs(seedForA - seedForB))
}

func (a *Card) compareOnPoints(b *Card) int {
	pointsForA := int(a.points())
	pointsForB := int(b.points())
	return pointsForA - pointsForB
}

func (a *Card) compareOnNumber(b *Card) int {
	numberForA := int(a.number)
	numberForB := int(b.number)
	return numberForA - numberForB
}
