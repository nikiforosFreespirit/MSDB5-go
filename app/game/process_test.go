package game

import (
	"testing"

	"github.com/mcaci/msdb5/app/phase"

	"github.com/mcaci/msdb5/dom/card"
)

func testGameSetup(withSide bool) *Game {
	gameTest := testGame(withSide)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.51", playerChannel)
	gameTest.Join("127.0.0.52", playerChannel)
	gameTest.Join("127.0.0.53", playerChannel)
	gameTest.Join("127.0.0.54", playerChannel)
	gameTest.Join("127.0.0.55", playerChannel)
	if withSide {
		gameTest.side.Clear()
		gameTest.side.Add(card.ID(31))
	}
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Draw(func() card.ID { return card.ID(2*i + 5) })
		if i > 1 {
			pl.Fold()
		}
	}
	return gameTest
}

func testGamePlay(gameTest *Game) {
	gameTest.Process("Join#A", "127.0.0.51")
	gameTest.Process("Join#B", "127.0.0.52")
	gameTest.Process("Join#C", "127.0.0.53")
	gameTest.Process("Join#D", "127.0.0.54")
	gameTest.Process("Join#E", "127.0.0.55")
	gameTest.Process("Auction#80", "127.0.0.51")
	gameTest.Process("Auction#79", "127.0.0.52")
	if gameTest.withSide {
		gameTest.Process("Exchange#5#Coin", "127.0.0.51")
		gameTest.Process("Exchange#0#Coin", "127.0.0.51")
	}
	gameTest.Process("Companion#7#Coin", "127.0.0.51")

	if gameTest.withSide {
		gameTest.Process("Card#1#Cudgel", "127.0.0.51")
	} else {
		gameTest.Process("Card#5#Coin", "127.0.0.51")
	}
	gameTest.Process("Card#7#Coin", "127.0.0.52")
	gameTest.Process("Card#9#Coin", "127.0.0.53")
	gameTest.Process("Card#1#Cup", "127.0.0.54")
	gameTest.Process("Card#3#Cup", "127.0.0.55")
}

func TestCompletedGameReturningScoreInfoWithSide(t *testing.T) {
	gameTest := testGameSetup(true)
	testGamePlay(gameTest)
	if gameTest.phase != phase.End {
		t.Fatal("Expecting transition to end game and scoring")
	}
}

func TestCompletedGameReturningScoreInfoWithNoSide(t *testing.T) {
	gameTest := testGameSetup(false)
	testGamePlay(gameTest)
	if gameTest.phase != phase.End {
		t.Fatal("Expecting transition to end game and scoring")
	}
}
