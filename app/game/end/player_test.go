package end

import (
	"strconv"
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type lastPlayerTestStr struct {
	currentPlayer *player.Player
	players       team.Players
	briscola      card.Item
}

func (lp lastPlayerTestStr) CurrentPlayer() *player.Player { return lp.currentPlayer }
func (lp lastPlayerTestStr) Cards() *set.Cards             { return lp.players[2].Hand() }
func (lp lastPlayerTestStr) Briscola() card.Item           { return lp.briscola }

func fakeGameSetup() *lastPlayerTestStr {
	currentPlayer := player.New()
	gameTest := lastPlayerTestStr{
		currentPlayer: currentPlayer,
		players:       team.Players{currentPlayer, player.New(), player.New(), player.New(), player.New()},
		briscola:      *card.MustID(1),
	}
	for i, pl := range gameTest.players {
		pl.RegisterAs(strconv.Itoa(i))
		pl.Hand().Add(*card.MustID(uint8(2*i + 5)))
	}
	return &gameTest
}

func TestCompletedGameReturningScoreInfoWithSide(t *testing.T) {
	gameTest := fakeGameSetup()
	lastPl := LastPlayer(*set.NewMust(1), gameTest.players, gameTest.players[0])
	if lastPl.Name() != "0" {
		t.Fatalf("Last player should be: %v", lastPl)
	}
}
