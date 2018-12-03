package bottles

// NewBottleNumber creates a BottleNumber from the number entered
func NewBottleNumber(n int) BottleNumber {
	switch n {
	case 0:
		return bottleNumber0{}
	case 1:
		return bottleNumber1{}
	default:
		return bottleNumberAll{n}
	}
}

// BottleNumber contains methods for BottleNumber
type BottleNumber interface {
	successor() int
	action() string
	quantity() string
	pronoun() string
	container() string
	String() string
}
