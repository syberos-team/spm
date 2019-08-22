package core


//返回码
const (
	CODE_SUCCESS string = "SUCCESS"
	CODE_ERROR string = "ERROR"
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
	Data *InfoResponseData	`json:"data"`
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

//LastVersionRequest 查询新版本
type LastVersionResponse struct {
	BaseResponse
	//返回数据
	Data LastVersionResponseData	`json:"data"`
}
//LastVersionResponseData 查询版本返回数据
type LastVersionResponseData struct {
	Version string 	`json:"version"`
}

//DownloadSpmRequest 下载spm请求参数
type DownloadSpmRequest struct {
	Version string 	`json:"version"`
}