package httpclient

import (
	"github.com/astroshot/going/pkg/helper"
)

func PostJSON(method *string, URL *string, header map[string]string, reqBody interface{}, resBody interface{}) {
	jsonBytes := helper.ToJSONBytes(reqBody)
	if header == nil {
		header = make(map[string]string)
		header["Content-Type"] = "application/json"
	}
	resBytes := Request(method, URL, header, jsonBytes)
	helper.BytesToStruct(*resBytes, resBody)
}
