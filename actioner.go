package bottles

import "fmt"

type actioner interface {
	action() string
}

type actioner0 struct{}

func (a actioner0) action() string {
	return "Go to the store and buy some more, "
}

type actionerDefault struct {
	pronouner pronouner
}

func (a actionerDefault) action() string {
	return fmt.Sprintf("Take %s down and pass it around, ", a.pronouner.pronoun())
}

type pronouner interface {
	pronoun() string
}

type pronouner1 struct{}

func (p pronouner1) pronoun() string {
	return "it"
}

type pronounerDefault struct{}

func (p pronounerDefault) pronoun() string {
	return "one"
}
