package declare

//TODO event相关的代码生成待实现

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
