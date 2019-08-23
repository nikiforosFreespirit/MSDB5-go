package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/team"
)

// PlInfo struct
type PlInfo struct {
	phase              phase.ID
	briscolaCard       card.Item
	players            team.Players
	playedCards        *set.Cards
	side, roundOngoing bool
	fromInput          string
}

func NewPlInfo(ph phase.ID, pls team.Players, plCards *set.Cards,
	briscola card.Item, isSide, isRound bool, frInput string) *PlInfo {
	return &PlInfo{phase: ph, players: pls, playedCards: plCards, fromInput: frInput,
		briscolaCard: briscola, side: isSide, roundOngoing: isRound}
}

func (nx PlInfo) Briscola() card.Item     { return nx.briscolaCard }
func (nx PlInfo) IsSideUsed() bool        { return nx.side }
func (nx PlInfo) IsRoundOngoing() bool    { return nx.roundOngoing }
func (nx PlInfo) Phase() phase.ID         { return nx.phase }
func (nx PlInfo) PlayedCards() *set.Cards { return nx.playedCards }
func (nx PlInfo) Players() team.Players   { return nx.players }
func (nx PlInfo) FromInput() string       { return nx.fromInput }