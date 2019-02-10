package bottles

import "fmt"

// newBottleNumber creates a new BottleNumber based on the number passed in.
func newBottleNumber(n int) bottleNumber {
	switch n {
	case 0:
		return bottleNumber{successorer0{}, actioner0{}, quantityer0{}, containererDefault{}}
	case 1:
		return bottleNumber{successorerDefault{n}, actionerDefault{pronouner1{}}, quantityerDefault{n}, containerer1{}}
	default:
		return bottleNumber{successorerDefault{n}, actionerDefault{pronounerDefault{}}, quantityerDefault{n}, containererDefault{}}
	}
}

// bottleNumber contains methods to describe the current number of bottles.
type bottleNumber struct {
	successorer successorer
	actioner    actioner
	quantityer  quantityer
	containerer containerer
}

func (b bottleNumber) successor() bottleNumber {
	return b.successorer.successor()
}

func (b bottleNumber) action() string {
	return b.actioner.action()
}

func (b bottleNumber) quantity() string {
	return b.quantityer.quantity()
}

func (b bottleNumber) container() string {
	return b.containerer.container()
}

// String method is used by fmt to get the toString representation of the given type.
func (b bottleNumber) String() string {
	return fmt.Sprintf("%s %s",
		b.quantity(),
		b.container())
}
