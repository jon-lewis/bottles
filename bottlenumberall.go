package bottles

import "fmt"

type bottleNumberAll struct {
	number int
}

func (n bottleNumberAll) successor() int {
	return n.number - 1
}

func (n bottleNumberAll) action() string {
	return fmt.Sprintf("Take %s down and pass it around, ", n.pronoun())
}

func (n bottleNumberAll) quantity() string {
	return fmt.Sprintf("%d", n.number)
}

func (n bottleNumberAll) pronoun() string {
	return "one"
}

func (n bottleNumberAll) container() string {
	return "bottles"
}

// String method is used by fmt to get the string representation of the given type.
func (n bottleNumberAll) String() string {
	return fmt.Sprintf("%s %s",
		n.quantity(),
		n.container())
}
