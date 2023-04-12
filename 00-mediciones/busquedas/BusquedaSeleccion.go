package busquedas

func Selection(items []int) []int {
	last := len(items) - 1
	for i := 0; i < last; i++ {
		aMin := items[i]
		iMin := i
		for j := i + 1; j < len(items); j++ {

			if items[j] < aMin {
				aMin = items[j]
				iMin = j
			}
		}
		items[i], items[iMin] = aMin, items[i]
	}
	return items
}
