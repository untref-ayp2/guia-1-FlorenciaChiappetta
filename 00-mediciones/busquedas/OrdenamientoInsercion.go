package busquedas

func Insertion(items []int) []int {

	for i := 1; i < len(items); i++ {

		value := items[i]
		j := i - 1
		for j >= 0 && items[j] > value {
			items[j+1] = items[j]
			j = j - 1
		}
		items[j+1] = value
	}
	return items
}
