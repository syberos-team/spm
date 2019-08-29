package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spm/core/log"
	"spm/core/util"
	"strings"
)


func getJson(rsp *http.Response, data *interface{}) error{
	result, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	defer util.CloseQuietly(rsp.Body)
	if log.IsDebug() {
		log.Debug("response.body:", string(result))
	}
	err = json.Unmarshal(result, data)
	if err != nil {
		return err
	}
	return nil
}

func GetJSON(url string, data *interface{}) error {
	log.Debug("request get:", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = getJson(resp, data)
	return err
}

func PostJSON(url, params string, data *interface{}) error{
	log.Debug("request post:", url, "params:", params)
	resp, err := http.Post(url, "application/json", strings.NewReader(params))
	if err!=nil {
		return err
	}
	err = getJson(resp, data)
	return err
}


func GetDownload(url string) (header *http.Header, data []byte, err error){
	log.Debug("request get:", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	header = &resp.Header
	defer util.CloseQuietly(resp.Body)
	data, err = ioutil.ReadAll(resp.Body)
	return
}
