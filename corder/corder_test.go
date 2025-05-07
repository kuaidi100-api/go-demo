package corder

import (
	"fmt"
	"testing"
	"time"
)

func TestCityOrder(t *testing.T) {

	Create()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Cancel()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Price()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Detail()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
