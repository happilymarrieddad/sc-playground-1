package utils

func Int64ArrToInterfaceArr(vals ...int64) (arr []interface{}) {
	for _, val := range vals {
		arr = append(arr, val)
	}

	return arr
}
