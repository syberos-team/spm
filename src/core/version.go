package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const SEP string = "."

type CheckVersionResponse struct {
	version string
}

//Version 版本号
type Version struct {
	//主版本号
	Major int
	//次要版本号
	Minor int
	//修订版本号
	Revision int
}

func (v *Version) String() string{
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Revision)
}

//CheckVersion 检查是否有新版本，存在新版本返回true，否则返回false
func (v *Version) CheckVersion() (bool, error){
	var data interface{}
	data = &CheckVersionResponse{}
	err := Get(API_CHECK_VERSION + "?version=" + v.String(), &data)
	if err != nil {
		return false, err
	}
	respData := data.(CheckVersionResponse)

	remoteVersion, err := Parse(respData.version)
	if err !=nil {
		return false, err
	}
	if remoteVersion.Major < v.Major {
		return false, nil
	}
	if remoteVersion.Major > v.Major {
		return true, nil
	}
	if remoteVersion.Minor < v.Minor {
		return false, nil
	}
	if remoteVersion.Minor > v.Minor {
		return true, nil
	}
	if remoteVersion.Revision < v.Revision {
		return false, nil
	}
	if remoteVersion.Revision > v.Revision {
		return true, nil
	}
	return false, nil
}

func Upgrade() error{
	return nil
}


//Parse 解析版本号字符串
func Parse(ver string) (*Version, error){
	verNums := strings.Split(ver, SEP)
	if verNums==nil || len(verNums) != 3 {
		return nil, errors.New("version number format error")
	}
	major, err := strconv.Atoi(verNums[0])
	if err!=nil {
		return nil, err
	}
	minor, err := strconv.Atoi(verNums[1])
	if err!=nil {
		return nil, err
	}
	revision, err := strconv.Atoi(verNums[1])
	if err!=nil {
		return nil, err
	}
	return &Version{major, minor, revision}, nil
}


