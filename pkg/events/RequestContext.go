package events

import (
	"fmt"
)

type RequestContext struct {
	ApiId     string `json:"apiId"`
	RequestId string `json:"requestId"`
	Stage     string `json:"stage"`
	SourceIp  string `json:"sourceIp"`
}

func (rc RequestContext) String() string {
	return fmt.Sprintf(`RequestContext{
                                 apiId='%s',
                                 requestId='%s',
                                 stage='%s',
                                 sourceIp='%s',
                               }`, rc.ApiId, rc.RequestId, rc.Stage, rc.SourceIp)
}
