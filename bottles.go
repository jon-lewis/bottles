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
	switch number {
	case 0:
		return "No more bottles of beer on the wall, " +
			"no more bottles of beer.\n" +
			"Go to the store and buy some more, " +
			"99 bottles of beer on the wall.\n"
	default:
		return fmt.Sprintf("%d %s of beer on the wall, "+
			"%d %s of beer.\n"+
			"Take %s down and pass it around, "+
			"%s %s of beer on the wall.\n",
			number, container(number), number, container(number), pronoun(number), quantity(number-1), container(number-1))
	}
}

func quantity(number int) string {
	switch number {
	case 0:
		return "no more"
	default:
		return fmt.Sprintf("%d", number)
	}
}

func pronoun(number int) string {
	switch number {
	case 1:
		return "it"
	default:
		return "one"
	}
}

func container(number int) string {
	switch number {
	case 1:
		return "bottle"
	default:
		return "bottles"
	}
}
