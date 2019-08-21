package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Struct2Params 将结构体转为请求参数map
func Struct2Params(t interface{}) *map[string][]string{
	m := make(map[string]string)
	data, _ := json.Marshal(t)
	_ = json.Unmarshal(data, m)

	result := make(map[string][]string)
	for k,v := range m {
		value := []string{v}
		result[k] = value
	}
	return &result
}

//Struct2Map 将结构体转为map
func Struct2Map(t interface{}) *map[string]interface{}{
	m := make(map[string]interface{})
	data, _ := json.Marshal(t)
	_ = json.Unmarshal(data, m)
	return &m
}

//Prompt 控制台输入问询，可传入默认值
func Prompt(prompt string, defaultValue string) chan string{
	replyChannel := make(chan string, 1)
	if defaultValue == "" {
		fmt.Print(prompt, " ")
	}else{
		fmt.Printf("%s [%s] ", prompt, defaultValue)
	}
	in := bufio.NewReader(os.Stdin)
	answer, _ := in.ReadString('\n')

	answer = strings.TrimSpace(answer)

	if len(answer) > 0 {
		replyChannel <- answer
	} else {
		replyChannel <- defaultValue
	}
	return replyChannel
}

//ToPrettyJSON 生成json字符串并格式化
func ToPrettyJSON(v interface{}) ([]byte, error){
	return json.MarshalIndent(v, "", "\t")
}

//LoadJsonFile 加载json文件内容转成struct
func LoadJsonFile(filePath string, data *interface{}) error{
	if !IsExists(filePath) {
		return errors.New("is not exists: " + filePath)
	}
	file, err := os.Open(filePath)
	if err!=nil {
		return err
	}
	defer CloseQuietly(file)
	bytes, err := ioutil.ReadAll(file)

	err = json.Unmarshal(bytes, data)
	if err!=nil {
		return err
	}
	return nil
}

//LoadTextFile 加载文本类型的文件
func LoadTextFile(filePath string) (string, error){
	if !IsExists(filePath) {
		return "", errors.New("is not exists: " + filePath)
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err!=nil {
		return "", err
	}
	return string(bytes), nil
}

//WriteStruct 将结构体转换成json格式字符串并写入文件
func WriteStruct(filePath string, data interface{}) error{
	bytes, err := ToPrettyJSON(data)
	if err!=nil {
		return err
	}
	return ioutil.WriteFile(filePath, bytes, os.FileMode(0666))
}

//ParsePackageInfo 解析包信息，拆分成包名和版本号
func ParsePackageInfo(pkgInfo string) (packageName, version string){
	packageInfo := strings.Split(pkgInfo, "@")
	if packageInfo==nil || len(packageInfo)==0{
		return
	}
	packageInfoLen := len(packageInfo)
	if packageInfoLen==1{
		packageName = packageInfo[0]
	}else if packageInfoLen==2 {
		packageName = packageInfo[0]
		version = packageInfo[1]
	}
	return
}