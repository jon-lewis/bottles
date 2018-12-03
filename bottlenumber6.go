package bottles

import "fmt"

type bottleNumber6 struct{}

func (n bottleNumber6) successor() int {
	return 5
}

func (n bottleNumber6) action() string {
	return fmt.Sprintf("Take %s down and pass it around, ", n.pronoun())
}

func (n bottleNumber6) quantity() string {
	return fmt.Sprintf("%d", 6)
}

func (n bottleNumber6) pronoun() string {
	return "one"
}

func (n bottleNumber6) container() string {
	return "bottles"
}

// String method is used by fmt to get the string representation of the given type.
func (n bottleNumber6) String() string {
	return fmt.Sprintf("%s %s",
		n.quantity(),
		n.container())
}
