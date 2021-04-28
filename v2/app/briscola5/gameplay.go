package briscola5

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola5/auction"
	"github.com/mcaci/msdb5/v2/app/briscola5/companion"
	"github.com/mcaci/msdb5/v2/app/briscola5/end"
	"github.com/mcaci/msdb5/v2/app/briscola5/exchange"
	"github.com/mcaci/msdb5/v2/app/briscola5/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Start(g *Game) {
	// distribute cards to players
	distributeCards(g)

	// auction phase
	aucInf := auction.Run(g.players)
	g.auctionScore = aucInf.Score
	g.players.SetCaller(aucInf.Caller)

	// card exchange phase
	if g.opts.WithSide {
		exchange.Run(struct {
			Hand, Side *set.Cards
		}{
			Hand: g.players.Caller().Hand(),
			Side: &g.side.Cards,
		})
	}

	// companion choice phase
	cmpInf := companion.Run(
		struct {
			ID      uint8
			Players team.Players
		}{
			ID:      aucInf.Caller,
			Players: briscola5.ToGeneralPlayers(g.players),
		},
	)
	g.briscolaCard = cmpInf.Briscola
	g.players.SetCaller(cmpInf.Companion)

	// play phase
	plInfo := play.Run(struct {
		Players      briscola5.Players
		BriscolaCard briscola.Card
	}{
		Players:      g.players,
		BriscolaCard: cmpInf.Briscola,
	})

	// end phase
	end.Run(struct {
		PlayedCards  briscola.PlayedCards
		Players      team.Players
		BriscolaCard briscola.Card
		Side         briscola5.Side
	}{
		PlayedCards:  plInfo.OnBoard,
		Players:      briscola5.ToGeneralPlayers(g.players),
		BriscolaCard: cmpInf.Briscola,
		Side:         g.side,
	})
}

func Score(g *Game) string {
	t1, t2 := briscola5.ToGeneralPlayers(g.players).Part(briscola5.IsInCallers(&g.players))
	return fmt.Sprintf("[%s: %d], [%s: %d]",
		"Caller team", briscola.Score(team.CommonPile(t1)),
		"Non Caller team", briscola.Score(team.CommonPile(t2)))
}

func distributeCards(g *Game) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if g.opts.WithSide && i >= set.DeckSize-5 {
			g.side.Add(d.Top())
			continue
		}
		g.players.Player(i % 5).Hand().Add(d.Top())
	}
}
