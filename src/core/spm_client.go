package core

import (
	"core/conf"
	"core/util"
	"fmt"
	"strings"
)

type SpmClient interface {
	Publish(req *PublishRequest) (*PublishResponse, error)
	GetDependency(req *DependencyRequest) (*DependencyResponse, error)
	Search(req *SearchRequest) (*SearchResponse, error)
	Info(req *InfoRequest) (*InfoResponse, error)

}

type BaseResponse struct {
	code string
	msg string
}

//PublishRequest 推送接口请求参数
type PublishRequest struct {
	//TODO
}

//PublishResponse 推送接口返回信息
type PublishResponse struct {
	BaseResponse
	//TODO
}

//DependencyRequest 依赖接口请求参数
type DependencyRequest struct {
	//TODO
}

//DependencyResponse 依赖接口返回信息
type DependencyResponse struct {
	BaseResponse
	//TODO
}

//SearchRequest 查询接口请求参数
type SearchRequest struct {
	//TODO
}

//SearchResponse 查询接口返回信息
type SearchResponse struct {
	BaseResponse
	//TODO
}

//InfoRequest 详情接口请求参数
type InfoRequest struct {
	//TODO
}

//InfoResponse 详情接口返回信息
type InfoResponse struct {
	BaseResponse
	//TODO
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
	err := s.sendGet(s.getUrl(API_SEARCH), &params, &result)
	return rsp, err
}

func (s *spmClient) Info(req *InfoRequest) (*InfoResponse, error) {
	rsp := &InfoResponse{}
	var result interface{} = rsp

	var params interface{} = req
	err := s.sendGet(s.getUrl(API_INFO), &params, &result)
	return rsp, err
}

//getUrl 拼接url
func (s *spmClient) getUrl(path string) string{
	url := strings.TrimSuffix(s.url, "/")
	if !strings.HasPrefix(path, "/") {
		url += "/"
	}
	return url + path
}

func (s *spmClient) sendPost(url string, params *interface{}, result *interface{}) error{
	p := util.Struct2Params(*params)
	return Post(url, p, result)
}

func (s *spmClient) sendGet(url string, params *interface{}, result *interface{}) error{
	builder := &strings.Builder{}
	builder.WriteString(url)
	args := util.Struct2Map(params)
	if len(*args) > 0 {
		builder.WriteString("?")
	}
	for k,v := range *args {
		builder.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	return Get(builder.String()[:builder.Len()-1], result)
}

//NewSpmClient 创建一个新的SpmClient客户端
func NewSpmClient() SpmClient{
	return &spmClient{url: conf.Config.Url}
}





