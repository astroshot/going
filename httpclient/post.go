package httpclient

import (
	"context"

	"github.com/astroshot/going/helper"
	"github.com/astroshot/going/mime"
)

// PostJSON do HTTP Post method, with JSON body
func PostJSON(ctx context.Context, URL *string, query map[string]string, header map[string]string, reqBody interface{}, resBody interface{}) {
	reqUrl := helper.MakeSortedURL(URL, query)
	jsonBytes := helper.ToJSONBytes(reqBody)
	if header == nil {
		header = make(map[string]string)
		header[mime.ContentType] = mime.AppJSON
	}

	var method = mime.Post
	resBytes := Request(ctx, reqUrl, &method, header, jsonBytes)
	helper.BytesToStruct(*resBytes, resBody)
}

// PostForm do HTTP Post method, with form body
func PostForm(ctx context.Context, URL *string, query map[string]string, header map[string]string, reqBody map[string]string, resBody interface{}) {
	reqUrl := helper.MakeSortedURL(URL, query)
	formBody := helper.MakeSortedForm(reqBody)
	formBytes := []byte(*formBody)
	if header == nil {
		header = make(map[string]string)
	}

	header[mime.ContentType] = mime.AppXForm
	var method = mime.Post
	resBytes := Request(ctx, reqUrl, &method, header, &formBytes)
	helper.BytesToStruct(*resBytes, resBody)
}
