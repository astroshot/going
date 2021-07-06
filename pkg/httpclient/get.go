package httpclient

import (
	"context"

	"github.com/astroshot/going/pkg/helper"
	"github.com/astroshot/going/pkg/mime"
)

// Get does HTTP GET request
func Get(ctx context.Context, URL *string, query map[string]string, header map[string]string, reqBody *[]byte, resBody interface{}) {
	var method = mime.Get
	resBytes := Request(ctx, &method, URL, header, reqBody)
	helper.BytesToStruct(*resBytes, resBody)
}
