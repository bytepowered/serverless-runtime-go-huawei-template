package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/runtime/pkg"
	"huaweicloud.com/runtime/pkg/events"
)

func onHandleRequest(payload []byte, ctx pkg.RuntimeContext) (interface{}, error) {
	var event events.TriggerEvent
	err := json.Unmarshal(payload, &event)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", event.String())
	resp := events.TriggerResponse{
		Body: event.String(),
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}
	return resp, nil
}

func main() {
	pkg.Register(onHandleRequest)
}
