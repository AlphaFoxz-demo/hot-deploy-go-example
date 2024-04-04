package declare

//go:generate go run ../gen_event_listener.go

import (
	"time"
)

type CheckedInEvent struct {
	Plate
	Time time.Time
}

type CheckInFailedEvent struct {
	Plate
	Time time.Time
}

type CheckedOutEvent struct {
	Plate
	Time time.Time
}

type CheckOutFailedEvent struct {
	Plate
	Time    time.Time
	Message string
}

type PaidEvent struct {
	Plate
	Amount  int
	PayTime time.Time
}
