package helper

import "encoding/json"

// ToJSONBytes dumps struct to byte array
func ToJSONBytes(v interface{}) *[]byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return &bytes
}

// ToJSON dumps struct to JSON string
func ToJSON(v interface{}) *string {
	bytes := ToJSONBytes(v)
	s := string(*bytes)
	return &s
}

// ToStruct convert JSON string to struct
func ToStruct(jsonStr *string, v interface{}) {
	var body []byte = []byte(*jsonStr)
	BytesToStruct(body, v)
}

// BytesToStruct converts byte arrary to struct
func BytesToStruct(body []byte, v interface{}) {
	if err := json.Unmarshal(body, v); err != nil {
		panic(err)
	}
}
