package helper

func IsNilString(str *string) bool {
	var nilString string
	return *str == nilString
}

func IsNilBool(b *bool) bool {
	var nilBool bool
	return *b == nilBool
}
