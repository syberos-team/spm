package core

import (
	"core/conf"
	"core/util"
	"encoding/json"
	"fmt"
	"strings"
)

//Package 包信息
type Package struct {
	//包名
	Name string `json:"name"`
	//描述
	Description string `json:"description"`
}
//Author 作者信息
type Author struct {
	//姓名
	Name string 	`json:"name"`
	//邮箱
	Email string	`json:"email"`
	//描述
	Description string `json:"description"`
}
//Repository 仓库信息
type Repository struct {
	//仓库url
	Url string 		`json:"url"`
}

//返回码
const (
	CODE_SUCCESS string = "SUCCESS"
	CODE_ERROR string = "ERROR"
)

//BaseResponse 响应返回通用信息
type BaseResponse struct {
	Code string		`json:"code"`
	Msg string		`json:"msg"`
}

//PublishRequest 推送接口请求参数
type PublishRequest struct {
	Package Package		`json:"package"`
	Author Author		`json:"author"`
	Repository Repository	`json:"repository"`
	Version string		`json:"version"`
	Dependencies []string	`json:"dependencies"`
	PriFilename	string 	`json:"priFilename"`
	Force	string		`json:"force"`
}

//PublishResponse 推送接口返回信息
type PublishResponse struct {
	BaseResponse
}

//DependencyRequest 依赖接口请求参数
type DependencyRequest struct {
	//git仓库url
	GitUrl string
}

//DependencyResponse 依赖接口返回信息
type DependencyResponse struct {
	BaseResponse
	//TODO
}

//SearchRequest 查询接口请求参数
type SearchRequest struct {
	//包名
	PackageName string	`json:"packageName"`
}

//SearchResponse 查询接口返回信息
type SearchResponse struct {
	BaseResponse
	//返回数据
	Data []*SearchResponseData		`json:"data"`
}
//SearchResponseData 查询接口数据
type SearchResponseData struct {
	//包名
	Name string 	`json:"name"`
	//描述
	Description string 	`json:"description"`
}

//InfoRequest 详情接口请求参数
type InfoRequest struct {
	//包名
	PackageName string 	`json:"packageName"`
	//版本号
	Version string 		`json:"version"`
}

//InfoResponse 详情接口返回信息
type InfoResponse struct {
	BaseResponse
	//返回数据
	Data InfoResponseData	`json:"data"`
}
//InfoResponseData 详情接口返回数据
type InfoResponseData struct {
	Package Package		`json:"package"`
	Author Author		`json:"author"`
	Repository Repository	`json:"repository"`
	Version string 		`json:"version"`
	Dependencies []string	`json:"dependencies"`
	PriFilename	string 	`json:"priFilename"`
}

type SpmClient interface {
	Publish(req *PublishRequest) (*PublishResponse, error)
	GetDependency(req *DependencyRequest) (*DependencyResponse, error)
	Search(req *SearchRequest) (*SearchResponse, error)
	Info(req *InfoRequest) (*InfoResponse, error)

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

//getUrl 拼接url
func (s *spmClient) getUrl(path string) string{
	url := strings.TrimSuffix(s.url, "/")
	if !strings.HasPrefix(path, "/") {
		url += "/"
	}
	return url + path
}

func (s *spmClient) sendPost(url string, params interface{}, result *interface{}) error{
	data, err := json.Marshal(params)
	if err!=nil {
		return err
	}
	return PostJSON(url, string(data), result)
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

