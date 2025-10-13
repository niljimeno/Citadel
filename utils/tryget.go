package utils

func TryGet(arr []string, it int) string {
	if len(arr) <= it {
		return ""
	} else {
		return arr[it]
	}
}
