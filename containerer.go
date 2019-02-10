package bottles

type containerer interface {
	container() string
}

type containerer1 struct{}

func (c containerer1) container() string {
	return "bottle"
}

type containererDefault struct{}

func (c containererDefault) container() string {
	return "bottles"
}
