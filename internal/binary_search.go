package internal

func BinarySearch[K Number](m []K, item K) K {
	var low K
	var high K
	var answer K
	low = 0
	high = len(m) - 1

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
		}

		if guess < item {
			low = middle + 1
		}

	}

	return answer
}
