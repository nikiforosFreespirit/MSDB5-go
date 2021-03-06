package player

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	if p := testP(""); !EmptyHanded(p) {
		t.Fatal("Player should not have cards at creation")
	}
}

func TestPlayerDrawsOneCard(t *testing.T) {
	p := testP("")
	p.Hand().Add(*card.MustID(1))
	plPredicate := IsCardInHand(*card.MustID(1))
	if !plPredicate(p) {
		t.Fatalf("Expecting player to have drawn %v", *card.MustID(1))
	}
}
