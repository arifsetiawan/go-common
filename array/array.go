package array

// IndexOfString is
func IndexOfString(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// DeleteElementString is
func DeleteElementString(data []string, element string) ([]string, int) {
	i := IndexOfString(data, element)
	if i == -1 {
		return data, i
	}

	data = append(data[:i], data[i+1:]...)
	return data, i
}

// AddElementString is
func AddElementString(data []string, element string) ([]string, int) {
	i := IndexOfString(data, element)
	if i == -1 {
		data = append(data, element)
	}

	return data, i
}

// DeleteElementStringArray is
func DeleteElementStringArray(data []string, elements []string) ([]string, int) {
	found := 0
	for _, element := range elements {
		i := IndexOfString(data, element)
		if i != -1 {
			found++
			data = append(data[:i], data[i+1:]...)
		}
	}

	return data, found
}

// AddElementStringArray is
func AddElementStringArray(data []string, elements []string) ([]string, int) {
	notFound := 0
	for _, element := range elements {
		foundIndex := IndexOfString(data, element)
		if foundIndex == -1 {
			notFound++
			data = append(data, element)
		}
	}

	return data, notFound
}
