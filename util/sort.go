package util

import "unsafe"

func QuickSort2(arrInter interface{}, compare func(x, y uintptr) bool) {
	arr := *(*[]uintptr)(unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&arrInter))[1]))
	if len(arr) <= 1 {
		return
	}
	i, j := 0, len(arr)-1
	flag := true
	for i < j {
		if !compare(arr[i], arr[j]) {
			arr[i], arr[j] = arr[j], arr[i]
			flag = !flag
		}
		if !flag {
			i++
		} else {
			j --
		}
	}
	if i > 1 {
		QuickSort2(arr[:i], compare)
	}
	if i+1 < len(arr)-2 {
		QuickSort2(arr[i+1:], compare)
	}
}
