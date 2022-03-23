package internal

func Mergesort[K Number](input []K) []K {
	if len(input) < 2 {
		return input
	}

	return fold(Mergesort(input[:int(len(input))/2]), Mergesort(input[int(len(input))/2:]))
}

func fold[K Number](listA, listB []K) []K {
	var i K
	var j K

	sortedList := []K{}

	for i < len(listA) && j < len(listB) {
		if listA[int(i)] < listB[int(j)] {
			sortedList = append(sortedList, listA[int(i)])
			i++
			continue
		}

		sortedList = append(sortedList, listB[int(j)])
		j++

	}

	//Left over elements
	for ; i < len(listA); i++ {
		sortedList = append(sortedList, listA[int(i)])
	}
	for ; j < len(listB); j++ {
		sortedList = append(sortedList, listB[int(j)])
	}

	return sortedList
}
