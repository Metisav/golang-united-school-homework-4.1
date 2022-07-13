package reverse_string

func ReverseString(input string) (output string) {
	inputRunes := []rune(input)
	length := len(inputRunes)
	outputRunes := make([]rune, length)

	for i, rune := range inputRunes {
		outputRunes[length-i-1] = rune
	}

	output = string(outputRunes)
	return output
}
