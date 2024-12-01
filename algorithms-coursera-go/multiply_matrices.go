package main

func multiplyMatrices(a, b [][]int) [][]int {
	aRows, aCols := len(a), len(a[0])
	bRows, bCols := len(b), len(b[0])

	var result [][]int

	if aCols != bRows {
		return nil
	}

	for i := 0; i < aRows; i++ {
		for j := 0; j < bCols; j++ {
			for k := 0; k < aCols; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result
}
