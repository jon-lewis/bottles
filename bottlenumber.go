package bottles

import "fmt"

// BottleNumber contains methods to describe the current number of bottles.
type BottleNumber struct {
	number int
}

func (bn BottleNumber) successor() int {
	switch bn.number {
	case 0:
		return 99
	default:
		return bn.number - 1
	}
}

func (bn BottleNumber) action() string {
	switch bn.number {
	case 0:
		return "Go to the store and buy some more, "
	default:
		return fmt.Sprintf("Take %s down and pass it around, ", bn.pronoun())
	}
}

func (bn BottleNumber) quantity() string {
	switch bn.number {
	case 0:
		return "no more"
	default:
		return fmt.Sprintf("%d", bn.number)
	}
}

func (bn BottleNumber) pronoun() string {
	switch bn.number {
	case 1:
		return "it"
	default:
		return "one"
	}
}

func (bn BottleNumber) container() string {
	switch bn.number {
	case 1:
		return "bottle"
	default:
		return "bottles"
	}
}
