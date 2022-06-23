package main

func main() {
	arr := []int{3, 1, 2, 3, 3}
	var result []int

	if len(arr) > 1 {
		for i := 0; i < len(arr)-1; i++ {
			var flag = true
			for n := i + 1; n < len(arr); n++ {
				if arr[i] == arr[n] {
					flag = false
				}
			}
			if flag {
				result = append(result, arr[i])
			}
		}
		result = append(result, arr[len(arr)-1])
	} else {
		result = arr
	}
}
