package bottles

import (
	"fmt"
	"strconv"
)

// Bottle is used to construct the verses in the 99 bottles song.
type Bottle struct{}

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

	return song
}
