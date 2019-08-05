package util

import (
	"bufio"
	"encoding/json"
	"fmt"
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