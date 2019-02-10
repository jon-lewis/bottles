package bottles

type actioner interface{}

type actioner0 struct{}

type pronouner interface{}

type pronouner1 struct{}

func (p pronouner1) pronoun() string {
	return "it"
}

type pronounerDefault struct{}

func (p pronounerDefault) pronoun() string {
	return "one"
}
