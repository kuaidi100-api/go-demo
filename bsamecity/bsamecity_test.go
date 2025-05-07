package bsamecity

import (
	"fmt"
	"testing"
	"time"
)

func TestBsamecity(t *testing.T) {
	Order()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Price()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Cancel()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	AddFee()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	PreCancel()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}
