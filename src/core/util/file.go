package util

import (
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

