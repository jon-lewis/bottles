package bottles

import "fmt"

// BottleNumberFactory contains methods for creating BottleNumbers.
type BottleNumberFactory struct{}

// NewBottleNumber creates a new BottleNumber based on the number passed in.
func (f BottleNumberFactory) NewBottleNumber(n int) BottleNumber {
	switch n {
	case 0:
		return BottleNumber{n, successorer0{}, actioner0{}, quantityer0{}, containererDefault{}}
	case 1:
		return BottleNumber{n, successorerDefault{n}, actionerDefault{pronouner1{}}, quantityerDefault{n}, containerer1{}}
	default:
		return BottleNumber{n, successorerDefault{n}, actionerDefault{pronounerDefault{}}, quantityerDefault{n}, containererDefault{}}
	}
}

// BottleNumber contains methods to describe the current number of bottles.
type BottleNumber struct {
	number      int
	successorer successorer
	actioner    actioner
	quantityer  quantityer
	containerer containerer
}

func (b BottleNumber) successor() int {
	return b.successorer.successor()
}

func (b BottleNumber) action() string {
	return b.actioner.action()
}

func (b BottleNumber) quantity() string {
	return b.quantityer.quantity()
}

func (b BottleNumber) container() string {
	return b.containerer.container()
}

// String method is used by fmt to get the string representation of the given type.
func (b BottleNumber) String() string {
	return fmt.Sprintf("%s %s",
		b.quantity(),
		b.container())
}
