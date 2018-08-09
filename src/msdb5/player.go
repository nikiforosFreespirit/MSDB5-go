package msdb5

import "container/list"

// Player interface
type Player interface {
	Init()
	Draw(d Deck) CardPtr
	Has(c CardPtr) bool
	Hasnt(c CardPtr) bool
}

// ConcretePlayer type
type ConcretePlayer struct {
	cards *list.List
}

// Init func
func (player *ConcretePlayer) Init() {
	player.cards = list.New()
}

// Draw func
func (player *ConcretePlayer) Draw(d Deck) CardPtr {
	c := d.RemoveTop()
	player.cards.PushFront(c)
	return c
}

// Has func
func (player *ConcretePlayer) Has(c CardPtr) bool {
	cardFound := false
	for e := player.cards.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

// Hasnt func
func (player *ConcretePlayer) Hasnt(c CardPtr) bool {
	return !player.Has(c)
}
