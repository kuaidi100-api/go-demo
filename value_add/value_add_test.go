package value_add

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	AddressResolution()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	AutoNumber()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	BackOrder()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	DetOCR()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	EstimatePrice()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	EstimateTime()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	InterceptOrder()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Reachable()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	SmsSend()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
