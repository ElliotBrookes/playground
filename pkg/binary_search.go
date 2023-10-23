package pkg

func BinarySearch(wantedValue int) bool {
	orderedData := [50]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

	lo := 0
	hi := 50

	for lo < hi {
		m := lo + ((hi - lo) / 2)
		v := orderedData[m]

		if wantedValue == v {
			return true
		} else if v > wantedValue {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}
