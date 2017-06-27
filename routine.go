package plz

import (
	"github.com/v2pro/plz/routine"
)

func Go(oneOff func()) error {
	return routine.Go(oneOff)
}

func GoLongRunning(longRunning func()) error {
	return routine.GoLongRunning(longRunning)
}