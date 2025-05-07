package border

import (
	"fmt"
	"testing"
	"time"
)

func TestBorder(t *testing.T) {

	Price()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Create()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Detail()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Cancel()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}
