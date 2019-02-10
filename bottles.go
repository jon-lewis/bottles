package bottles

import (
	"fmt"

	"github.com/jon-lewis/words"
)

// Bottle is used to construct the verses in the 99 bottles song.
type Bottle struct{}

func (bottle Bottle) song() string {
	return bottle.verses(99, 0)
}

func (bottle Bottle) verses(x int, y int) string {
	song := bottle.verse(x)

	for i := x - 1; i >= y; i-- {
		song += "\n" + bottle.verse(i)
	}

	return song
}

// Verse method sings the 99 bottles song
func (bottle Bottle) verse(number int) string {
	bottleNumber := BottleNumberFactory{}.NewBottleNumber(number)
	nextBottleNumber := BottleNumberFactory{}.NewBottleNumber(bottleNumber.successor())

	return fmt.Sprintf("%s of beer on the wall, "+
		"%s of beer.\n"+
		"%s"+
		"%s of beer on the wall.\n",
		words.Capitalize(bottleNumber.String()),
		bottleNumber,
		bottleNumber.action(),
		nextBottleNumber)
}
