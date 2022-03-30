package internal

func Quicksort[K Number](input []K) []K {
	if len(input) < 2 {
		return input
	}
	// Pivot in the middle
	pivot := input[int(len(input)/2)]
	lessThan := []K{}
	moreThan := []K{}
	
	for x, item := range input {
		// Don't evaluate the Pivot
		if x == int(len(input)/2) {
			continue
		}

		if item < pivot {
			lessThan = append(lessThan, item)
			continue
		}

		moreThan = append(moreThan, item)
	}

	return append(Quicksort(lessThan), append([]K{pivot}, Quicksort(moreThan)...)...)

}
