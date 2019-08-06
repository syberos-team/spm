package core

import (
	"core/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)


func getJson(rsp *http.Response, data *interface{}) error{
	result, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	defer util.CloseQuietly(rsp.Body)
	err = json.Unmarshal(result, data)
	if err != nil {
		return err
	}
	return nil
}

func Get(url string, data *interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = getJson(resp, data)
	return err
}

func Post(url string, params *map[string][]string, data *interface{}) error{
	resp, err := http.PostForm(url, *params)
	if err != nil {
		return err
	}
	err = getJson(resp, data)
	return err
}

func PostJSON(url, params string, data *interface{}) error{
	resp, err := http.Post(url, "application/json", strings.NewReader(params))
	if err!=nil {
		return err
	}
	err = getJson(resp, data)
	return err
}


func GetDownload(url string, writer io.Writer) (size int64, err error){
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer util.CloseQuietly(resp.Body)
	size, err = io.Copy(writer, resp.Body)
	return
}

func PostDownload(url string, params *map[string][]string, writer io.Writer) (size int64, err error){
	resp, err := http.PostForm(url, *params)
	if err != nil {
		return
	}
	defer util.CloseQuietly(resp.Body)
	size, err = io.Copy(writer, resp.Body)
	return
}

