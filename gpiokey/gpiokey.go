package gpiokey

import (
	"github.com/e-asphyx/gpio"
	"github.com/e-asphyx/input"
	"runtime"
	"time"
)

type GpioKey struct {
	Ch      chan input.Event
	trigger gpio.PinTrigger
	active  ActiveLevel
	value   int
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
	value, err := pin.Read()
	if err != nil {
		return nil, err
	}

	trigger, err := pin.Trigger(gpio.EdgeBoth)
	if err != nil {
		return nil, err
	}

	key = &GpioKey{
		Ch:      ch,
		trigger: trigger,
		active:  active,
		value:   value,
		key:     keycode,
	}
	go key.serve()
	runtime.SetFinalizer(key, (*GpioKey).Close)

	return key, nil
}

func (key *GpioKey) serve() {
	timer := time.NewTimer(debounceInterval)
	var debounce bool = false

	for {
		select {
		case <-timer.C:
			debounce = false

		case val, ok := <-key.trigger.Ch():
			if !ok {
				return
			}

			if val != key.value {
				key.value = val

				if !debounce {
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

					debounce = true
					timer.Reset(debounceInterval)
				}
			}
		}
	}
}

func (key *GpioKey) Close() error {
	return key.trigger.Close()
}
