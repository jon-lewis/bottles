package bottles

import "fmt"

type quantityer interface {
	quantity() string
}

type quantityer0 struct{}

func (q quantityer0) quantity() string {
	return "no more"
}

type quantityerDefault struct {
	number int
}

func (q quantityerDefault) quantity() string {
	return fmt.Sprintf("%d", q.number)
}
