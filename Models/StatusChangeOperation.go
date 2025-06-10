package Models

type StatusChangeOperation int

const (
	Upgrade   StatusChangeOperation = iota
	Downgrade                       = 1
)
