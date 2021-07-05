package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Request(url *string, method *string, header map[string]string, body *[]byte) *[]byte {
	client := &http.Client{}

	reqBody := bytes.NewBuffer(*body)
	req, err := http.NewRequest(*method, *url, reqBody)
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
