package util

//StringSliceRemove 删除字符串类型切片中指定下标的元素
func StringSliceRemove(slice []string, index int) []string{
	return append(slice[:index], slice[index+1:]...)
}
