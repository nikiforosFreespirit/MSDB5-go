package card

import (
	"errors"
	"strconv"
)

// ByID func
func ByID(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

func (id ID) idToNumber() uint8 {
	return id.toZeroBased()%10 + 1
}

func (id ID) idToSeed() Seed {
	return Seed(id.toZeroBased() / 10)
}

func (id ID) toZeroBased() uint8 {
	return uint8(id) - 1
}

func (id ID) String() string {
	return "(" + strconv.Itoa(int(id.idToNumber())) + " of " + id.idToSeed().String() + ")"
}

