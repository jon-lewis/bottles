package bottles

type successorer interface {
	successor() int
}

type successorer0 struct{}

func (s successorer0) successor() int {
	return 99
}

type successorerDefault struct {
	number int
}

func (s successorerDefault) successor() int {
	return s.number - 1
}
