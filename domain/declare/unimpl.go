package declare

import (
	"container/list"
)

//需要外部实现的功能

type ParkingRepo interface {
	FindById(id string) ParkingAgg
}
type EventQueue interface {
	Enqueue(event CheckedInEvent)
	GetItems() list.List
}
