package international

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	ApiCall()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	PickUp()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	CancelPickUp()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	InternationalAddressResolution()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
