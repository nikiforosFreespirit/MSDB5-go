package phase

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
)

type valueProvider interface{ Value() string }

type cardProvider interface{ Card() (card.ID, error) }

type Data struct {
	name string

	toFold bool
	score  auction.Score

	card    card.ID
	plIdx   uint8
	pl      *player.Player
	cardErr error
}

func (d Data) Name() string { return d.name }

func (d Data) ToFold() bool         { return d.toFold }
func (d Data) Score() auction.Score { return d.score }

func (d Data) Card() card.ID      { return d.card }
func (d Data) Index() uint8       { return d.plIdx }
func (d Data) Pl() *player.Player { return d.pl }
func (d Data) CardErr() error     { return d.cardErr }
