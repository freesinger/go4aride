package concurrent

import (
	"github.com/freesinger/go4aride/concurrent"
	"testing"
)

func TestSendAndPrint(t *testing.T) {
	concurrent.SentAndReceive()
}

func TestKickBall(t *testing.T) {
	concurrent.Football()
}