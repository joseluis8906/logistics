package sales

import (
	"fmt"
)

type (
	Currency float64
)

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", c)
}
