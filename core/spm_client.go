package core

import (
	"encoding/json"
	"fmt"
	"io"
	"spm/core/conf"
	"spm/core/util"
	"strings"
)


type SpmClient interface {
	Publish(req *PublishRequest) (*PublishResponse, error)
	GetDependency(req *DependencyRequest) (*DependencyResponse, error)
	Search(req *SearchRequest) (*SearchResponse, error)
	Info(req *InfoRequest) (*InfoResponse, error)
	LastVersion() (*LastVersionResponse, error)
	DownloadSpm(version string, w *io.Writer) error
}

//spmClient spm客户端结构体
type spmClient struct {
	url string
}

func (s *spmClient) Publish(req *PublishRequest) (*PublishResponse, error){
	rsp := &PublishResponse{}
	var result interface{} = rsp

	var params interface{} = req
	err := s.sendPost(s.getUrl(API_PUBLISH), &params, &result)
	return rsp, err
}

func (s *spmClient) GetDependency(req *DependencyRequest) (*DependencyResponse, error) {
	rsp := &DependencyResponse{}
	var result interface{} = rsp
	var params interface{} = req
	err := s.sendGet(s.getUrl(API_GET_DEPENDENCY), &params, &result)
	return rsp, err
}

func (s *spmClient) Search(req *SearchRequest) (*SearchResponse, error) {
	rsp := &SearchResponse{}
	var result interface{} = rsp

	var params interface{} = req
	err := s.sendPost(s.getUrl(API_SEARCH), &params, &result)
	return rsp, err
}

func (s *spmClient) Info(req *InfoRequest) (*InfoResponse, error) {
	rsp := &InfoResponse{}
	var result interface{} = rsp

	var params interface{} = req
	err := s.sendPost(s.getUrl(API_INFO), &params, &result)
	return rsp, err
}

func (s *spmClient) LastVersion() (*LastVersionResponse, error){
	rsp := &LastVersionResponse{}
	var result interface{} = rsp
	err := s.sendGet(s.getUrl(API_LAST_VERSION), nil , &result)
	return rsp, err
}

func (s *spmClient) DownloadSpm(version string, w *io.Writer) error {
	req := &DownloadSpmRequest{
		Version: version,
	}
	var params interface{} = req
	_, err := GetDownload(s.getUrlWithParams(API_DOWNLOAD_SPM, &params), w)
	if err!=nil {
		return err
	}
	return nil
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
	url := s.getUrl(path)
	if params==nil {
		return url
	}
	builder := &strings.Builder{}
	builder.WriteString(url)
	args := util.Struct2Map(params)
	if len(*args) > 0 {
		builder.WriteString("?")
	}
	for k,v := range *args {
		builder.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	return builder.String()[:builder.Len()-1]
}

func (s *spmClient) sendPost(url string, params interface{}, result *interface{}) error{
	data, err := json.Marshal(params)
	if err!=nil {
		return err
	}
	return PostJSON(url, string(data), result)
}

func (s *spmClient) sendGet(url string, params *interface{}, result *interface{}) error{
	if params==nil {
		return Get(url, result)
	}

	return Get(url, result)
}

//NewSpmClient 创建一个新的SpmClient客户端
func NewSpmClient() SpmClient{
	return &spmClient{url: conf.Config.Url}
}

