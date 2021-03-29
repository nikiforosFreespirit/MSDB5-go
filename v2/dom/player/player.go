package player

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
)

// Player struct
type Player struct {
	name       string
	hand, pile set.Cards
}

// New func
func New() *Player {
	return new(Player)
}

// Name func
func (player Player) Name() string { return player.name }

// Hand func
func (player *Player) Hand() *set.Cards { return &player.hand }

// Pile func
func (player *Player) Pile() *set.Cards { return &player.pile }

// RegisterAs func
func (player *Player) RegisterAs(name string) { player.name = name }

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v)",
		player.name, player.hand, player.pile)
}
