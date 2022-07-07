package reverse_string

func ReverseString(input string) (output string) {

	in_runes := []rune(input)
	out_runes := make([]rune, len(input))
	for k, v := range in_runes {
		out_runes[len(in_runes)-k-1] = v
	}
	output = string(out_runes)
	return
}
