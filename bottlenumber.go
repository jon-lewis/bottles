package bottles

import "fmt"

// BottleNumber contains methods to describe the current number of bottles.
type BottleNumber struct {
	number int
}

func (b BottleNumber) successor() int {
	switch b.number {
	case 0:
		return 99
	default:
		return b.number - 1
	}
}

func (b BottleNumber) action() string {
	switch b.number {
	case 0:
		return "Go to the store and buy some more, "
	default:
		return fmt.Sprintf("Take %s down and pass it around, ", b.pronoun())
	}
}

func (b BottleNumber) quantity() string {
	switch b.number {
	case 0:
		return "no more"
	default:
		return fmt.Sprintf("%d", b.number)
	}
}

func (b BottleNumber) pronoun() string {
	switch b.number {
	case 1:
		return "it"
	default:
		return "one"
	}
}

func (b BottleNumber) container() string {
	switch b.number {
	case 1:
		return "bottle"
	default:
		return "bottles"
	}
}

// String method is used by fmt to get the string representation of the given type.
func (b BottleNumber) String() string {
	return fmt.Sprintf("%s %s",
		b.quantity(),
		b.container())
}
