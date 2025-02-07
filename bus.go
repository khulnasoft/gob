package gob

import "github.com/khulnasoft/go-pulsebus"

type BusConstructor func(Config) *pulsebus.Bus

var _ BusConstructor = newBus

func newBus(_ Config) *pulsebus.Bus {
	return pulsebus.NewBus()
}
