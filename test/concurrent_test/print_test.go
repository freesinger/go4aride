package concurrent_test

import (
	"github.com/freesinger/go4aride/concurrent"
	"testing"
)

func TestPrint(t *testing.T) {
	concurrent.Print()
}

// Deadlock
func TestSleep(t *testing.T) {
	concurrent.Sleep()
}