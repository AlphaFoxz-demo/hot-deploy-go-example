package declare

import (
	"time"
)

type CheckInCommand struct {
	repo_         ParkingRepo
	Plate         Plate
	CheckedInTime time.Time
}

func (command *CheckInCommand) Handle() bool {
	agg := command.repo_.FindById(command.Plate.Number)
	result := agg.HandleCheckIn(*command)
	return result
}

type CheckOutCommand struct {
	repo_         ParkingRepo
	Plate         Plate
	CheckedInTime time.Time
}

//
//type NotifyPayCommand struct {
//	repo_   *ParkingRepo
//	Plate   Plate
//	Amount  int
//	PayTime time.Time
//}
