package bottles

import "fmt"

type bottleNumber1 struct{}

func (n bottleNumber1) successor() int {
	return 0
}

func (n bottleNumber1) action() string {
	return fmt.Sprintf("Take %s down and pass it around, ", n.pronoun())
}

func (n bottleNumber1) quantity() string {
	return fmt.Sprintf("%d", 1)
}

func (n bottleNumber1) pronoun() string {
	return "it"
}

func (n bottleNumber1) container() string {
	return "bottle"
}

// String method is used by fmt to get the string representation of the given type.
func (n bottleNumber1) String() string {
	return fmt.Sprintf("%s %s",
		n.quantity(),
		n.container())
}
