package end

import (
	"container/list"
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

var messageSink = func(p *player.Player, msg string) {}

type fakeGame struct {
	comp, call *player.Player
	players    team.Players
}

func fakePlayer(name, host string, ids ...card.ID) *player.Player {
	p := player.New()
	p.RegisterAs(name)
	p.Join(host)
	for _, id := range ids {
		p.Draw(func() card.ID { return id })
	}
	return p
}

func newTestGame3Cards() fakeGame {
	f := fakeGame{}
	t := team.Players{}
	f.comp = fakePlayer("A", "127.0.0.51", 10, 5, 5)
	t.Add(f.comp)
	t.Add(fakePlayer("B", "127.0.0.52", 5, 5, 5))
	f.call = fakePlayer("C", "127.0.0.53", 1, 3, 6)
	t.Add(f.call)
	t.Add(fakePlayer("D", "127.0.0.54", 5, 5, 5))
	t.Add(fakePlayer("E", "127.0.0.55", 5, 5, 5))
	f.players = t
	return f
}

func newTestGame3CardsNotEndable() fakeGame {
	f := fakeGame{}
	t := team.Players{}
	f.comp = fakePlayer("A", "127.0.0.51", 5, 5, 5)
	t.Add(f.comp)
	t.Add(fakePlayer("B", "127.0.0.52", 10, 5, 5))
	f.call = fakePlayer("C", "127.0.0.53", 1, 3, 6)
	t.Add(f.call)
	t.Add(fakePlayer("D", "127.0.0.54", 5, 5, 5))
	t.Add(fakePlayer("E", "127.0.0.55", 5, 5, 5))
	f.players = t
	return f
}

func newTestGameNoCards() fakeGame {
	f := fakeGame{}
	t := team.Players{}
	f.comp = fakePlayer("A", "127.0.0.51")
	t.Add(f.comp)
	t.Add(fakePlayer("B", "127.0.0.52"))
	f.call = fakePlayer("C", "127.0.0.53")
	t.Add(f.call)
	t.Add(fakePlayer("D", "127.0.0.54"))
	t.Add(fakePlayer("E", "127.0.0.55"))
	f.players = t
	return f
}

func (g fakeGame) Caller() *player.Player    { return g.call }
func (g fakeGame) Companion() *player.Player { return g.comp }
func (g fakeGame) Players() team.Players     { return g.players }
func (g fakeGame) LastPlaying() *list.List   { return list.New() }
func (g fakeGame) Briscola() card.Seed       { return card.Coin }

func TestTrueCheckWithCardLeft(t *testing.T) {
	gameTest := newTestGame3Cards()
	err := Check(gameTest, messageSink)
	if !err {
		t.Fatal(err)
	}
}

func TestFalseCheckWithCardLeft(t *testing.T) {
	gameTest := newTestGame3CardsNotEndable()
	err := Check(gameTest, messageSink)
	if err {
		t.Fatal(err)
	}
}

func TestTrueCheckWithNoCardLeft(t *testing.T) {
	gameTest := newTestGameNoCards()
	err := Check(gameTest, messageSink)
	if !err {
		t.Fatal(err)
	}
}

func TestProcessWithNoErr(t *testing.T) {
	gameTest := newTestGameNoCards()
	err := Process(gameTest, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}