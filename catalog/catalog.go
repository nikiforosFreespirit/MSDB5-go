// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p := messageKeyToIndex[key]
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"it": &dictionary{index: itIndex, data: itData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"(Turn of: %s, Companion is: %s, Played cards: %+v, Auction score: %d, Phase: %d)": 10,
	"Action %s not valid":                      3,
	"Error: %+v\n":                             0,
	"Expecting player %s to play":              1,
	"Game: %+v":                                7,
	"Phase is not %d but %d":                   2,
	"Player: %+v\n":                            4,
	"Side deck section: %+v\n":                 6,
	"Side deck: %+v\n":                         5,
	"The end - %s team has all briscola cards": 8,
	"The end - Callers: %d; Others: %d":        9,
}

var enIndex = []uint32{ // 12 elements
	0x00000000, 0x00000013, 0x00000032, 0x0000004f,
	0x00000066, 0x0000007a, 0x00000091, 0x000000b0,
	0x000000bd, 0x000000e9, 0x00000111, 0x00000171,
} // Size: 72 bytes

const enData string = "" + // Size: 369 bytes
	"\x04\x00\x01\x0a\x0e\x02Error: %+[1]v\x02Expecting player %[1]s to play" +
	"\x02Phase is not %[1]d but %[2]d\x02Action %[1]s not valid\x04\x00\x01" +
	"\x0a\x0f\x02Player: %+[1]v\x04\x00\x01\x0a\x12\x02Side deck: %+[1]v\x04" +
	"\x00\x01\x0a\x1a\x02Side deck section: %+[1]v\x02Game: %+[1]v\x02The end" +
	" - %[1]s team has all briscola cards\x02The end - Callers: %[1]d; Others" +
	": %[2]d\x02(Turn of: %[1]s, Companion is: %[2]s, Played cards: %+[3]v, A" +
	"uction score: %[4]d, Phase: %[5]d)"

var itIndex = []uint32{ // 12 elements
	0x00000000, 0x00000014, 0x00000037, 0x00000055,
	0x0000006d, 0x00000084, 0x00000097, 0x000000b6,
	0x000000d3, 0x000000fa, 0x00000128, 0x00000189,
} // Size: 72 bytes

const itData string = "" + // Size: 393 bytes
	"\x04\x00\x01\x0a\x0f\x02Errore: %+[1]v\x02Mi aspetto che sia %[1]s a gio" +
	"care\x02La fase non e' %[1]d ma %[2]d\x02Azione %[1]s non valida\x04\x00" +
	"\x01\x0a\x12\x02Giocatore: %+[1]v\x04\x00\x01\x0a\x0e\x02Monte: %+[1]v" +
	"\x04\x00\x01\x0a\x1a\x02Sezione del monte: %+[1]v\x02Informazioni di gio" +
	"co %+[1]v\x02Fine - I %[1]s hanno tutte le briscole\x02Fine - Chiamanti:" +
	" %[1]d; Non-chiamanti: %[2]d\x02(Turno di: %[1]s, Compagno: %[2]s, Carte" +
	" giocate: %+[3]v, Puntegggio d'asta: %[4]d, Fase: %[5]d)"

	// Total table size 906 bytes (0KiB); checksum: E9633B83