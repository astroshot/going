package httpclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

// Request does basic HTTP Request
func Request(ctx context.Context, url *string, method *string, header map[string]string, body *[]byte) *[]byte {
	if ctx == nil {
		ctx = context.Background()
	}

	client := &http.Client{}
	reqBody := bytes.NewBuffer(*body)
	req, err := http.NewRequestWithContext(ctx, *method, *url, reqBody)
	if err != nil {
		panic(err)
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return &resBody
}
