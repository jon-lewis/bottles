package bottles

import (
	"fmt"
	"strconv"
)

// Bottle is used to construct the verses in the 99 bottles song.
type Bottle struct{}

func (b Bottle) verses(x int, y int) string {
	song := b.verse(x)

	for i := x - 1; i >= y; i-- {
		song += "\n" + b.verse(i)
	}

	return song
}

// Verse method sings the 99 bottles song
func (b Bottle) verse(x int) string {
	bottleDescription := "bottles"
	nextBottleDescription := "bottles"
	itOrOne := "one"
	startOfLastLine := strconv.Itoa(x - 1)

	if x-1 == 1 {
		nextBottleDescription = "bottle"
	} else if x-1 == 0 {
		itOrOne = "it"
		bottleDescription = "bottle"
		startOfLastLine = "no more"
	}

	song := fmt.Sprintf("%d %s of beer on the wall, "+
		"%d %s of beer.\n"+
		"Take %s down and pass it around, "+
		"%s %s of beer on the wall.\n",
		x,
		bottleDescription,
		x,
		bottleDescription,
		itOrOne,
		startOfLastLine,
		nextBottleDescription)

	if x == 0 {
		song = "No more bottles of beer on the wall, " +
			"no more bottles of beer.\n" +
			"Go to the store and buy some more, " +
			"99 bottles of beer on the wall.\n"
	}

	return song
}
