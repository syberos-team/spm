package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"spm/core/conf"
	"strconv"
	"strings"
	"syscall"
)

//当前版本号
const VERSION = "1.0.0"

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


//Version 版本号
type VersionManage struct {
	Version
	//最新的版本号，需要更新时此处存在
	lastVersion Version
}

//CheckVersion 检查是否有新版本，存在新版本返回true，否则返回false
func (v *VersionManage) CheckVersion() (bool, error){
	client := NewSpmClient()
	rsp, err := client.LastVersion(VERSION)
	if err!=nil {
		return false, err
	}
	if rsp.Code!=CODE_SUCCESS {
		return false, errors.New(rsp.Msg)
	}

	remoteVersion, err := ParseVersion(rsp.Data.Version)
	if err !=nil {
		return false, err
	}

	if remoteVersion.Major < v.Major {
		return false, nil
	}
	if remoteVersion.Major > v.Major {
		v.lastVersion = *remoteVersion
		return true, nil
	}
	if remoteVersion.Minor < v.Minor {
		return false, nil
	}
	if remoteVersion.Minor > v.Minor {
		v.lastVersion = *remoteVersion
		return true, nil
	}
	if remoteVersion.Revision < v.Revision {
		return false, nil
	}
	if remoteVersion.Revision > v.Revision {
		v.lastVersion = *remoteVersion
		return true, nil
	}
	return false, nil
}

func (v *VersionManage) Upgrade() error{
	fmt.Println("Ready to update to", v.lastVersion.String())
	client := NewSpmClient()

	filename := conf.FILENAME + ".tmp"
	filePath := path.Join(conf.Config.GetConfigDir(), filename)
	data, err := client.DownloadSpm(v.lastVersion.String())
	if err!=nil {
		return err
	}
	if err = ioutil.WriteFile(filePath, data, os.ModePerm); err!=nil {
		return err
	}
	spmPath, err := os.Executable()
	if err!=nil {
		return err
	}
	args := []string{filename, "copy", filePath, spmPath}
	env := os.Environ()
	if err = syscall.Exec(filePath, args, env); err!=nil {
		panic(err)
	}
	return nil
}

//ParseVersion 解析版本号字符串
func ParseVersion(ver string) (*Version, error){
	verNums := strings.Split(ver, ".")
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
	revision, err := strconv.Atoi(verNums[2])
	if err!=nil {
		return nil, err
	}
	return &Version{major, minor, revision}, nil
}

func NewVersionManage() VersionManage{
	v, _ := ParseVersion(VERSION)
	return VersionManage{Version:*v}
}


