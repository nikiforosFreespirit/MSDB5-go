package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type AuctionStruct struct {
	request, origin string
	players         playerset.Players
	board           *board.Board
}

func NewAuction(request, origin string, players playerset.Players, board *board.Board) Action {
	return &AuctionStruct{request, origin, players, board}
}
func (as AuctionStruct) Do(p *player.Player) error {
	data := strings.Split(as.request, "#")
	score := data[1]
	auction.CheckAndUpdate(score, p.Folded, p.Fold, as.board.AuctionScore, as.board.SetAuctionScore)
	return nil
}
func (as AuctionStruct) NextPlayer(playerInTurn uint8) uint8 {
	winnerIndex := playersRoundRobin(playerInTurn)
	for as.players[winnerIndex].Folded() {
		winnerIndex = playersRoundRobin(winnerIndex)
	}
	return winnerIndex
}
func (as AuctionStruct) NextPhase() game.Phase {
	var isFolded = func(p *player.Player) bool { return p.Folded() }
	if as.players.Count(isFolded) == 4 {
		if len(*as.board.SideDeck()) > 0 {
			return game.ExchangingCards
		}
		return game.ChosingCompanion
	}
	return game.InsideAuction
}
