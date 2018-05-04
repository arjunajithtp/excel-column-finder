package helpers

func Reverse(str string) string {
	if str != "" {
		return Reverse(str[1:]) + str[:1]
	}
	return ""
}
