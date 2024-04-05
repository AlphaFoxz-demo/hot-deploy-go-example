package declare

func (command *CheckInCommand) Init(repo_ ParkingRepo) (v *CheckInCommand) {
	command.repo_ = repo_
	return command
}
