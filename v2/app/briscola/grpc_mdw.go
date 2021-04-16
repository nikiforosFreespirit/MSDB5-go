package briscola

import (
	"context"
	"fmt"
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/pb"
)

func ScoreGrpc(g *Game) string {
	conn := pb.Conn()
	defer conn.Close()
	client := pb.NewBriscolaClient(conn)

	toPBCards := func(cards set.Cards) *pb.Cards {
		pbcards := make([]*pb.CardID, len(cards))
		for i := range pbcards {
			pbcards[i] = &pb.CardID{Id: uint32(cards[i].ToID())}
		}
		return &pb.Cards{Cards: pbcards}
	}

	s1, err := client.Score(context.Background(), toPBCards(*g.players.Players[0].Pile()))
	if err != nil {
		log.Println(err)
	}
	log.Println(s1.GetPoints())

	s2, err := client.Score(context.Background(), toPBCards(*g.players.Players[1].Pile()))
	if err != nil {
		log.Println(err)
	}
	log.Println(s2.GetPoints())

	return fmt.Sprintf("[%s: %d], [%s: %d]", "Caller team", s1.Points, "Non Caller team", s2.Points)
}
