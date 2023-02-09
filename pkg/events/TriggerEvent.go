package events

import (
	"encoding/base64"
	"fmt"
)

type TriggerEvent struct {
	IsBase64Encoded       bool              `json:"isBase64Encoded"`
	HttpMethod            string            `json:"httpMethod"`
	Path                  string            `json:"path"`
	Body                  string            `json:"body"`
	PathParameters        map[string]string `json:"pathParameters"`
	RequestContext        RequestContext    `json:"requestContext"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	UserData              string            `json:"user_data"`
}

func (e *TriggerEvent) GetRawBody() string {
	decoded, err := base64.StdEncoding.DecodeString(e.Body)
	if err != nil {
		return ""
	}
	return string(decoded)
}

func (e *TriggerEvent) String() string {
	return fmt.Sprintf(`TriggerEvent{
                                 isBase64Encoded=%v,
                                 httpMethod='%v',
                                 path='%v',
                                 body='%v',
                                 pathParamters=%v,
                                 requestContext=%v,
                                 headers=%v,
                                 queryStringParameters=%v,
                                 user_data=%+v,
                              }`, e.IsBase64Encoded, e.HttpMethod, e.Path, e.Body, e.PathParameters, e.RequestContext,
		e.Headers, e.QueryStringParameters, e.UserData)
}
