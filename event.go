package gob

import (
	"os"

	"github.com/khulnasoft/go-pulsebus"
)

const ExitEventType pulsebus.EventType = "gob-exit"

func ExitEvent(interrupt bool) pulsebus.Event {
	if interrupt {
		return pulsebus.Event{
			Type:  ExitEventType,
			Value: os.Interrupt,
		}
	}
	return pulsebus.Event{
		Type: ExitEventType,
	}
}
