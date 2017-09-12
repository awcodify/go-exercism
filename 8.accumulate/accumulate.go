package accumulate

const testVersion = 1

func Accumulate(input []string, converter func(string) string) []string {
	if input == nil || len(input) == 0 {
		return input
	}

	var mappedInput []string
	for _, element := range input {
		mappedInput = append(mappedInput, converter(element))
	}
	return mappedInput
}
