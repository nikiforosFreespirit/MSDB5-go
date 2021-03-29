package game

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

// Game struct
type Game struct {
	players      briscola5.Players
	cTeam        briscola5.Callerer
	briscolaCard card.Item
	side         set.Cards
	auctionScore auction.Score
	opts         *Options
}

type Options struct {
	WithSide bool
}

func NewGame(gOpts *Options) *Game { return &Game{opts: gOpts} }

// New func
func New() *Game { return &Game{} }

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.cTeam.Caller().Name(), g.cTeam.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}
