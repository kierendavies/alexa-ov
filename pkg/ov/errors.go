package ov

import (
	"fmt"
)

type NoArrivalsError struct {
	TimingPoint string
}

func (err *NoArrivalsError) Error() string {
	return fmt.Sprint("no arrivals found at timing point %s", err.TimingPoint)
}
