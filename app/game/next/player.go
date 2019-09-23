package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type nextPlayerInformer interface {
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	IsRoundOngoing() bool
	FromInput() string
}

// Player func
func Player(g nextPlayerInformer) *player.Player {
	numberOfPlayers := uint8(len(g.Players()))
	playersRoundRobin := func(playerIndex uint8) uint8 { return (playerIndex + 1) % numberOfPlayers }
	index, _ := g.Players().Find(player.MatchingHost(g.FromInput()))
	playerIndex := uint8(index)
	nextPlayer := playersRoundRobin(playerIndex)
	switch g.Phase() {
	case phase.InsideAuction:
		for player.Folded(g.Players()[nextPlayer]) {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.PlayingCards:
		if g.IsRoundOngoing() {
			break
		}
		winningCardIndex := indexOfWinningCard(*g.PlayedCards(), g.Briscola().Seed())
		nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
	}
	return g.Players()[nextPlayer]
}

func indexOfWinningCard(cardsOnTheTable set.Cards, b card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(base, other, b) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func winningCard(base, other card.Item, b card.Seed) card.Item {
	s := briscola.NewSorted(set.Cards{base, other}, &b)
	if &base == nil || s.Less(1, 0) {
		base = other
	}
	return base
}
