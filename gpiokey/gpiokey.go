package gpiokey

import (
	"github.com/e-asphyx/gpio"
	"github.com/e-asphyx/input"
	"time"
)

type GpioKey struct {
	Ch      chan input.Event
	trigger gpio.PinTrigger
	active  ActiveLevel
	key     input.Key
}

type ActiveLevel int

const (
	ActiveHigh ActiveLevel = iota
	ActiveLow
)

const debounceInterval = 5 * time.Millisecond

func (act ActiveLevel) Value(value int) int {
	if act == ActiveHigh {
		return value
	} else {
		return (^value) & 1
	}
}

func NewGpioKey(pin gpio.PinReadTrigger, active ActiveLevel, keycode input.Key) (key *GpioKey, err error) {
	ch := make(chan input.Event)
	return NewGpioKeyWithChannel(pin, active, keycode, ch)
}

func NewGpioKeyWithChannel(pin gpio.PinReadTrigger, active ActiveLevel, keycode input.Key, ch chan input.Event) (key *GpioKey, err error) {
	trigger, err := pin.TriggerWithDebounce(gpio.EdgeBoth, gpio.DefaultDebounceInterval)
	if err != nil {
		return nil, err
	}

	key = &GpioKey{
		Ch:      ch,
		trigger: trigger,
		active:  active,
		key:     keycode,
	}

	go func() {
		for val := range key.trigger.Ch() {
			var state input.KeyState
			if key.active.Value(val) == 1 {
				state = input.KeyPressed
			} else {
				state = input.KeyReleased
			}

			ts := time.Now()
			keyEvt := &input.KeyEvent{
				Time:  ts,
				Key:   key.key,
				State: state,
			}
			key.Ch <- keyEvt
			key.Ch <- &input.SynEvent{ts}
		}
	}()

	return key, nil
}

func (key *GpioKey) Close() error {
	return key.trigger.Close()
}

func (key *GpioKey) Active() ActiveLevel {
	return key.active
}

func (key *GpioKey) Key() input.Key {
	return key.key
}
