package bottles

import "fmt"

// BottleNumber contains methods to describe the current number of bottles.
type BottleNumber struct {
	number int
}

func (b BottleNumber) successor() int {
	switch b.number {
	case 0:
		return successorer0{}.successor()
	default:
		return successorerDefault{b.number}.successor()
	}
}

func (b BottleNumber) action() string {
	switch b.number {
	case 0:
		return actioner0{}.action()
	case 1:
		return actionerDefault{pronouner1{}}.action()
	default:
		return actionerDefault{pronounerDefault{}}.action()
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
