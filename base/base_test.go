package base

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	Query()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Poll()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	PollMap()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	MapTrack()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}
