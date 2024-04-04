package declare

func (command *CheckInCommand) Init(repo_ ParkingRepo) (v *CheckInCommand) {
	command.repo_ = repo_
	return command
}
func (command *CheckOutCommand) Init(repo_ ParkingRepo) (v *CheckOutCommand) {
	command.repo_ = repo_
	return command
}
