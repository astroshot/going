package helper

// BytesToStr converts byte array to string
func BytesToStr(bytes *[]byte) *string {
	if bytes == nil {
		return nil
	}

	if len(*bytes) < 1 {
		s := ""
		return &s
	}

	s := string(*bytes)
	return &s
}

// StrToBytes converts string to byte array
func StrToBytes(s *string) *[]byte {
	if s == nil {
		return nil
	}

	if len(*s) < 1 {
		bytes := make([]byte, 0)
		return &bytes
	}

	bytes := []byte(*s)
	return &bytes
}
