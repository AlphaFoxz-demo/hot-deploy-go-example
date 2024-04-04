package declare

import (
	"container/list"
)

type ParkingRepo interface {
	FindById(id string) ParkingAgg
}
type EventQueue interface {
	Enqueue(event CheckedInEvent)
	GetItems() list.List
}
