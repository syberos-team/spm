package core

import (
	"encoding/json"
	"errors"
	"reflect"
	"spm/core/conf"
	"strings"
)


type SpmClient interface {
	Publish(req *PublishRequest) (*PublishResponse, error)
	GetDependency(req *DependencyRequest) (*DependencyResponse, error)
	Search(req *SearchRequest) (*SearchResponse, error)
	Info(req *InfoRequest) (*InfoResponse, error)
	LastVersion(version string) (*LastVersionResponse, error)
	DownloadSpm(version string) ([]byte, error)
}

//spmClient spm客户端结构体
type spmClient struct {
	url string
}

func (s *spmClient) Publish(req *PublishRequest) (*PublishResponse, error){
	rsp := &PublishResponse{}
	var result interface{} = rsp
	err := s.sendPost(s.getUrl(conf.ApiPublish), req, &result)
	return rsp, err
}

func (s *spmClient) GetDependency(req *DependencyRequest) (*DependencyResponse, error) {
	rsp := &DependencyResponse{}
	var result interface{} = rsp
	err := s.sendGet(s.getUrl(conf.ApiGetDependency), req, &result)
	return rsp, err
}

func (s *spmClient) Search(req *SearchRequest) (*SearchResponse, error) {
	rsp := &SearchResponse{}
	var result interface{} = rsp
	err := s.sendGet(s.getUrl(conf.ApiSearch), req, &result)
	return rsp, err
}

func (s *spmClient) Info(req *InfoRequest) (*InfoResponse, error) {
	rsp := &InfoResponse{}
	var result interface{} = rsp
	err := s.sendGet(s.getUrl(conf.ApiInfo), req, &result)
	return rsp, err
}

func (s *spmClient) LastVersion(version string) (*LastVersionResponse, error){
	rsp := &LastVersionResponse{}
	var result interface{} = rsp

	req := &LastVersionRequest{Version:version}
	err := s.sendGet(s.getUrl(conf.ApiLastVersion), req, &result)
	return rsp, err
}

func (s *spmClient) DownloadSpm(version string) ([]byte, error) {
	req := &DownloadSpmRequest{
		Version: version,
	}
	header, data, err := GetDownload(s.getUrlWithParams(conf.ApiDownloadSpm, req))
	if err!=nil {
		return nil, err
	}
	contentType := header.Get("Content-Type")
	if strings.Contains(contentType, "application/json"){
		rsp := &BaseResponse{}
		err = json.Unmarshal(data, rsp)
		if err!=nil {
			return nil, err
		}
		return nil, errors.New(rsp.Msg)
	}

	return data, nil
}

//getUrl 拼接url
func (s *spmClient) getUrl(path string) string{
	url := strings.TrimSuffix(s.url, "/")
	if !strings.HasPrefix(path, "/") {
		url += "/"
	}
	return url + path
}

func (s *spmClient) getUrlWithParams(path string, params interface{}) string{
	if params==nil {
		return s.getUrl(path)
	}

	paramsValue := reflect.ValueOf(params)
	if paramsValue.Kind()==reflect.Ptr {
		paramsValue = paramsValue.Elem()
	}
	paramsType := paramsValue.Type()

	builder := strings.Builder{}

	for i := 0; i<paramsType.NumField(); i++ {
		fieldTag := paramsType.Field(i).Tag
		name := fieldTag.Get("form")
		if name=="" {
			name = fieldTag.Get("json")
		}
		value := paramsValue.Field(i).String()
		builder.WriteString(name)
		builder.WriteString("=")
		builder.WriteString(value)
		builder.WriteString("&")
	}
	urlParams := builder.String()[: builder.Len()-1]
	return s.getUrl(path) + "?" + urlParams
}

func (s *spmClient) sendPost(url string, params interface{}, result *interface{}) error{
	data, err := json.Marshal(params)
	if err!=nil {
		return err
	}
	return PostJSON(url, string(data), result)
}

//发送get请求，注意参数params必须是一个结构体
func (s *spmClient) sendGet(url string, params interface{}, result *interface{}) error{
	if params==nil {
		return GetJSON(url, result)
	}

	paramsValue := reflect.ValueOf(params)
	if paramsValue.Kind()==reflect.Ptr {
		paramsValue = paramsValue.Elem()
	}
	paramsType := paramsValue.Type()

	builder := strings.Builder{}

	for i := 0; i<paramsType.NumField(); i++ {
		fieldTag := paramsType.Field(i).Tag
		name := fieldTag.Get("form")
		if name=="" {
			name = fieldTag.Get("json")
		}
		value := paramsValue.Field(i).String()
		builder.WriteString(name)
		builder.WriteString("=")
		builder.WriteString(value)
		builder.WriteString("&")
	}
	urlParams := builder.String()[: builder.Len()-1]
	return GetJSON(url + "?" + urlParams, result)
}

//NewSpmClient 创建一个新的SpmClient客户端
func NewSpmClient() SpmClient{
	return &spmClient{url: conf.Config.Url}
}

