package bottles

import "fmt"

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
	verse := ""

	if number > 2 {
		verse = fmt.Sprintf("%d bottles of beer on the wall, "+
			"%d bottles of beer.\n"+
			"Take one down and pass it around, "+
			"%d bottles of beer on the wall.\n",
			number, number, number-1)
	} else if number == 2 {
		verse = "2 bottles of beer on the wall, " +
			"2 bottles of beer.\n" +
			"Take one down and pass it around, " +
			"1 bottle of beer on the wall.\n"
	} else if number == 1 {
		verse = "1 bottle of beer on the wall, " +
			"1 bottle of beer.\n" +
			"Take it down and pass it around, " +
			"no more bottles of beer on the wall.\n"
	} else if number == 0 {
		verse = "No more bottles of beer on the wall, " +
			"no more bottles of beer.\n" +
			"Go to the store and buy some more, " +
			"99 bottles of beer on the wall.\n"
	}

	return verse
}
