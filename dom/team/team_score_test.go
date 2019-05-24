package team

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

type mockScorer struct{}

func (mockScorer) Count(func(card.ID) uint8) uint8 {
	return 1
}

func TestTeam1(t *testing.T) {
	fakePlayer := new(mockScorer)
	if score1, _ := Score(fakePlayer, nil, fakePlayer); score1 != 1 {
		t.Fatal("Count string should contain the total of 1")
	}
}

func TestTeam2(t *testing.T) {
	if _, score2 := Score(nil, nil, new(mockScorer)); score2 != 1 {
		t.Fatal("Count string should contain the total of 1")
	}
}