package internal

func BinarySearch[K Number](m []K, item K) K {
	var low K = 0
	var high K = len(m) - 1
	var answer K

	for low < high {
		var middle K
		middle = low + high
		guess := m[int(middle)]

		if guess == item {
			answer = middle
			break
		}

		if guess > item {
			high = middle - 1
			continue
		}

		if guess < item {
			low = middle + 1
			continue
		}

	}

	return answer
}
