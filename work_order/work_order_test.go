package work_order

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	WorkOrderCreate()

	WorkOrderQuery()

	WorkOrderAddReply()

	WorkOrderReply()

	file, err := os.Open("C:\\***\\event.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	WorkOrderUpload(file)
}
