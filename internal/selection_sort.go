package internal


func SelectionSort[T Number](list []T) []T {
	inputLen := len(list)

	for i := 0; i < int(inputLen); i++ {
		minimumIndex := i

		for j := i; j < int(inputLen); j++ {
			if list[j] < list[minimumIndex] {
				minimumIndex = j
			}
		}

		list[i], list[minimumIndex] = list[minimumIndex], list[i]

	}
	return list
}
