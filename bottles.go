package bottles

import (
	"fmt"

	"github.com/jon-lewis/words"
)

// Bottle is used to construct the verses in the 99 bottles song.
type Bottle struct{}

func (b Bottle) song() string {
	return b.verses(99, 0)
}

func (b Bottle) verses(x int, y int) string {
	song := b.verse(x)

	for i := x - 1; i >= y; i-- {
		song += "\n" + b.verse(i)
	}

	return song
}

// Verse method sings the 99 bottles song
func (b Bottle) verse(number int) string {
	bn := BottleNumber{number}
	sbn := BottleNumber{bn.successor()}

	return fmt.Sprintf("%s %s of beer on the wall, "+
		"%s %s of beer.\n"+
		"%s"+
		"%s %s of beer on the wall.\n",
		words.Capitalize(bn.quantity()), bn.container(),
		bn.quantity(), bn.container(),
		bn.action(),
		sbn.quantity(), sbn.container())
}
