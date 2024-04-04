package declare

import (
	"time"
)

// Plate 车牌
type Plate struct {
	// 号码
	Number string
}

type ParkingAgg struct {
	EventQueue_ EventQueue
	Id          Plate
	CheckInTime time.Time
	LastPayTime time.Time
	TotalPaid   int
}

func (agg *ParkingAgg) HandleCheckIn(command CheckInCommand) bool {
	if agg.IsInPark() {
		return false
	}
	//TODO
	//agg.EventQueue_.Enqueue(CheckedInEvent{Plate: agg.Id, Time: command.CheckedInTime})
	agg.CheckInTime = command.CheckedInTime
	return true
}

func (agg *ParkingAgg) IsInPark() bool {
	return !agg.CheckInTime.IsZero()
}
