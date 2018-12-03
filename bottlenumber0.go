package bottles

import "fmt"

type bottleNumber0 struct{}

func (n bottleNumber0) successor() int {
	return 99
}

func (n bottleNumber0) action() string {
	return "Go to the store and buy some more, "
}

func (n bottleNumber0) quantity() string {
	return "no more"
}

func (n bottleNumber0) container() string {
	return "bottles"
}

func (n bottleNumber0) pronoun() string {
	return "it"
}

// String method is used by fmt to get the string representation of the given type.
func (n bottleNumber0) String() string {
	return fmt.Sprintf("%s %s",
		n.quantity(),
		n.container())
}
