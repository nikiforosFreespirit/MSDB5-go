package briscola

import (
	"errors"
	"log"

	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

const nPlayers = 2

type Players struct{ team.Players }

// NewPlayers creates new container for briscola5 players
func NewPlayers() *Players {
	players := make(team.Players, nPlayers)
	for i := range players {
		players[i] = player.New()
	}
	return &Players{Players: players}
}

func (pls *Players) Registration() func(string) error {
	var i int
	return func(s string) error {
		if i >= nPlayers {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, s)
		pls.At(i).RegisterAs(s)
		i++
		return nil
	}
}

func (pls *Players) At(i int) *player.Player       { return pls.Players[i] }
func (pls *Players) All(prd player.Predicate) bool { return pls.Players.All(prd) }
