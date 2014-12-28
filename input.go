package input

import "time"

type Event interface {
	Timestamp() time.Time
}

//go:generate stringer -type=Key
type Key int

//go:generate stringer -type=KeyState
type KeyState int

type KeyEvent struct {
	Time  time.Time
	Key   Key
	State KeyState
	Sym   rune
}

type SynEvent struct {
	Time time.Time
}

const (
	KeyReleased KeyState = iota
	KeyPressed
)

func (evt *KeyEvent) Timestamp() time.Time {
	return evt.Time
}

func (evt *SynEvent) Timestamp() time.Time {
	return evt.Time
}
