package util

import (
	"fmt"
	"io"
	"os"
)

func IsExists(p string) bool{
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//CloseQuietly 安静的调用Close()
func CloseQuietly(closer io.Closer){
	_ = closer.Close()
}

//Pwd 当前所在目录的路径
func Pwd() (string, error){
	if cwd, err := os.Getwd(); err != nil {
		return "", err
	}else{
		return cwd, err
	}
}


func CopyFile(src, dst string) (int64, error){
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer CloseQuietly(source)

	var destination *os.File
	if IsExists(dst) {
		destination, err = os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err!=nil {
			return 0, err
		}
	}else {
		destination, err = os.Create(dst)
		if err!=nil {
			return 0, err
		}
	}
	defer CloseQuietly(destination)
	return io.Copy(destination, source)
}
