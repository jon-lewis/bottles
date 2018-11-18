package bottles

import (
	"fmt"
	"strings"
	"unicode"
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
	return fmt.Sprintf("%s %s of beer on the wall, "+
		"%s %s of beer.\n"+
		"%s"+
		"%s %s of beer on the wall.\n",
		capitalize(quantity(number)), container(number),
		quantity(number), container(number),
		action(number),
		quantity(successor(number)), container(successor(number)))
}

func successor(number int) int {
	switch number {
	case 0:
		return 99
	default:
		return number - 1
	}
}

func action(number int) string {
	switch number {
	case 0:
		return "Go to the store and buy some more, "
	default:
		return fmt.Sprintf("Take %s down and pass it around, ", pronoun(number))
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

func capitalize(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if first {
				first = false
				return unicode.ToUpper(r)
			}
			return r
		},
		s)
}
