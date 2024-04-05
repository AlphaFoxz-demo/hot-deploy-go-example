package api

import (
	. "github.com/AlphaFoxz/hot-deploy-go-example/domain/declare"
	"time"
)

type api struct {
	repo_ ParkingRepo
}

func New(repo_ ParkingRepo) (v api) {
	return api{
		repo_: repo_,
	}
}
func (inst_ api) NewCheckInCommand(Plate Plate, CheckedInTime time.Time) (v CheckInCommand) {
	command := CheckInCommand{
		Plate:         Plate,
		CheckedInTime: CheckedInTime,
	}
	return *command.Init(inst_.repo_)
}
