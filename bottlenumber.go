package bottles

import "fmt"

// BottleNumber contains methods to describe the current number of bottles.
type BottleNumber struct {
	number int
}

func (bottleNumber BottleNumber) successor() int {
	switch bottleNumber.number {
	case 0:
		return 99
	default:
		return bottleNumber.number - 1
	}
}

func (bottleNumber BottleNumber) action() string {
	switch bottleNumber.number {
	case 0:
		return "Go to the store and buy some more, "
	default:
		return fmt.Sprintf("Take %s down and pass it around, ", bottleNumber.pronoun())
	}
}

func (bottleNumber BottleNumber) quantity() string {
	switch bottleNumber.number {
	case 0:
		return "no more"
	default:
		return fmt.Sprintf("%d", bottleNumber.number)
	}
}

func (bottleNumber BottleNumber) pronoun() string {
	switch bottleNumber.number {
	case 1:
		return "it"
	default:
		return "one"
	}
}

func (bottleNumber BottleNumber) container() string {
	switch bottleNumber.number {
	case 1:
		return "bottle"
	default:
		return "bottles"
	}
}
