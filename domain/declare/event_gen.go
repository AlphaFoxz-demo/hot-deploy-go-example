package declare

type CheckedInEventListener interface {
	onEvent(event CheckedInEvent)
}
type CheckInFailedEventListener interface {
	onEvent(event CheckInFailedEvent)
}
