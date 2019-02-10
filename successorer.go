package bottles

type successorer interface {
	successor() bottleNumber
}

type successorer0 struct{}

func (s successorer0) successor() bottleNumber {
	return newBottleNumber(99)
}

type successorerDefault struct {
	number int
}

func (s successorerDefault) successor() bottleNumber {
	return newBottleNumber(s.number - 1)
}
