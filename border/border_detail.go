package border

import (
	"encoding/json"
	"fmt"
	"time"
)

func Detail() {
	method := "detail"

	param := Param{
		TaskId: "9FC293CA417E431F33046E64F4C4EC20",
	}

	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	_, err := DoBorderRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
