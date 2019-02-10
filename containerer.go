package bottles

type containerer interface {
	container() string
}

type containerer0 struct{}

func (c containerer0) container() string {
	return "bottle"
}

type containererDefault struct{}

func (c containererDefault) container() string {
	return "bottles"
}
