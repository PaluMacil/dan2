package www

// MakeRows takes a slice of T and chunks it by the requested number of columns into a slice of rows of T
func MakeRows[T any](list []T, cols int) [][]T {
	rows := make([][]T, 0)
	if cols == 0 {
		return rows
	}
	var currentRow int
	for i, item := range list {
		// no remainder means start of a new row
		if i%cols == 0 {
			currentRow = i / cols
			row := make([]T, 0)
			rows = append(rows, row)
		}
		rows[currentRow] = append(rows[currentRow], item)
	}
	return rows
}
