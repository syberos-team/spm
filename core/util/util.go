package util

//StringSliceRemove 删除字符串类型切片中指定下标的元素
func StringSliceRemove(slice []string, index int) []string{
	return append(slice[:index], slice[index+1:]...)
}

//StringSlicePrepend 向字符串类型切片前端添加元素
func StringSlicePrepend(slice []string, elems... string) []string{
	newSlice := elems[:]
	return append(newSlice, slice...)
}

//SlicePrepend 向切片前端添加元素
func SlicePrepend(slice []interface{}, elems... interface{}) []interface{} {
	newSlice := elems[:]
	return append(newSlice, slice...)
}
