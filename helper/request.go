package helper

import (
	"fmt"
	"net/url"
	"sort"
)

func MakeForm(query map[string]string) *string {
	if query == nil || len(query) < 1 {
		return nil
	}

	queryParams := url.Values{}
	for k, v := range query {
		queryParams.Add(k, v)
	}
	s := queryParams.Encode()
	return &s
}

func MakeSortedForm(query map[string]string) *string {
	if query == nil || len(query) < 1 {
		return nil
	}

	keys := GetKeys(query)
	sort.Strings(*keys)
	queryParams := url.Values{}
	for _, k := range *keys {
		v := query[k]
		queryParams.Add(k, v)
	}
	s := queryParams.Encode()
	return &s
}

func formatURL(URL *string, queryParams *string) *string {
	var reqUrl string
	reqUrl = fmt.Sprintf("%s?%s", *URL, *queryParams)
	return &reqUrl
}

func MakeURL(URL *string, query map[string]string) *string {
	queryParams := MakeForm(query)
	if queryParams == nil {
		return URL
	}

	return formatURL(URL, queryParams)
}

func MakeSortedURL(URL *string, query map[string]string) *string {
	queryParams := MakeSortedForm(query)
	if queryParams == nil {
		return URL
	}

	return formatURL(URL, queryParams)
}
