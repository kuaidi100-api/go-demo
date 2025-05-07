package elec

import (
	"fmt"
	"testing"
	"time"
)

// 测试函数
func Test(t *testing.T) {
	Order()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	printOld()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	Cancel()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	custom()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	AuthThird()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	QueryBalance()
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
