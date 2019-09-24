package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin
	verifySortingWithBriscola(t, 1, 2, nil)
}

func TestScenarioWithTwoOfCoinLosing(t *testing.T) {
	// 2 and 3 of Coin
	verifySortingWithBriscola(t, 3, 2, nil)
}

func TestScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
	// 5 and 6 of Coin
	verifySortingWithBriscola(t, 6, 5, nil)
}

func TestScenarioWithTenOfCoinWinning(t *testing.T) {
	// 10 and 4 of Coin
	verifySortingWithBriscola(t, 10, 4, nil)
}

func TestPoints(t *testing.T) {
	if Points(*card.MustID(1)) != Points(*card.MustID(21)) {
		t.Fatal("unexpected result")
	}
}
