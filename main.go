package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/events"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
)

func onHandleRequest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var event events.APIGTriggerEvent
	err := json.Unmarshal(payload, &event)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", event.String())
	resp := events.APIGTriggerResponse{
		Body: event.String(),
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}
	return resp, nil
}

func main() {
	runtime.Register(onHandleRequest)
}
