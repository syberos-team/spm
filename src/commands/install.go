package commands

import (
	"bufio"
	"core/util"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)



type SpmJsonContent struct {
	Dependencies []string `json:"dependencies"`
}


const (
	Vendor = "vendor"
	VendorPri = "vendor.pri"
)

type InstallCommand struct {
	Command
	packageName string		//包名，通过参数获得
	version string		//版本号，通过参数获得，若参数中没有是，通过接口获取

	priFilename string 		//pri文件名，通过接口获取
}

func (i *InstallCommand) Description() string {
	return "Installs a new package"
}

func (i *InstallCommand) Run() error {
	spmJsonContent := &SpmJsonContent{}

	//检查spm.json文件中是否已存在待安装的包
	spmJsonFilePath := path.Join(pwd, SpmJsonFilename)
	if util.IsExists(spmJsonFilePath) {
		err := i.loadSpmJson(spmJsonFilePath, spmJsonContent)
		if err!=nil {
			return err
		}
		err = i.checkDependency(*spmJsonContent)
		if err!=nil {
			return err
		}
	}
	//查询待安装的包信息
	infoData, err := PackageInfo(i.packageName, i.version)
	if err!=nil {
		return err
	}
	i.version = infoData.Version
	i.priFilename = infoData.PriFilename
	//创建verdor目录
	vendorDirPath := path.Join(pwd, Vendor)
	if !util.IsExists(vendorDirPath) {
		err = os.MkdirAll(vendorDirPath, os.ModePerm)
		if err!=nil {
			return err
		}
	}
	//检查仓库地址
	repoUrl := infoData.Repository.Url
	if repoUrl=="" {
		return errors.New("no repository URL was obtained")
	}
	//下载源码
	subdir := path.Join(vendorDirPath, strings.ReplaceAll(infoData.Package.Name, ".", "/"))
	err = i.downloadFromGit(repoUrl, subdir)
	//更新spm.json
	err = i.updateSpmJson(spmJsonFilePath, spmJsonContent)
	if err!=nil {
		return err
	}
	//更新vendor.pri
	err = i.updateVendor(path.Join(vendorDirPath, VendorPri))
	if err!=nil {
		return err
	}
	return nil
}

func (i *InstallCommand) RegisterArgs(args ...string) {
	if args==nil || len(args)==0 {
		return
	}
	i.packageName, i.version = util.ParsePackageInfo(args[0])
}

func (i *InstallCommand) downloadFromGit(url, path string) error{
	err := util.GitClone(url, path)
	if err!=nil {
		return err
	}
	err = util.RemoveDotGit(path)
	if err!=nil {
		return err
	}
	return nil
}

func (i *InstallCommand) createVendorPriContent(vendorPriPath string) (string, error){
	var oldContent string
	if util.IsExists(vendorPriPath) {
		var err error
		oldContent, err = util.LoadTextFile(vendorPriPath)
		if err!=nil {
			return "", err
		}
	}

	var includePris []string

	includeNewPri := "include($$PWD/" +
		path.Join(strings.ReplaceAll(i.packageName, ".", "/"), i.priFilename) +
		")"

	oldContentBuf := bufio.NewReader(strings.NewReader(oldContent))
	for{
		line, err := oldContentBuf.ReadString('\n')
		if io.EOF==err {
			break
		}
		if err!=nil {
			return "", err
		}
		if strings.Contains(line, includeNewPri) {
			break
		}
		if strings.HasPrefix(line, "include(") {
			includePris = append(includePris, strings.TrimSpace(line))
		}
	}
	includePris = append(includePris, includeNewPri)

	newContent, err := util.TemplateToString(util.VendorTemplate, util.TemplateModel{IncludePris: includePris})
	if err!=nil {
		return "", err
	}
	return newContent, nil
}

//updateVendor 更新vendor.pri文件，不存在时会新建，若文件中已存在安装包的pri文件路径时，pri文件仍会重写，但内容不会变化
func (i *InstallCommand) updateVendor(vendorPriPath string) error{
	content, err := i.createVendorPriContent(vendorPriPath)
	if err!=nil {
		return err
	}
	return ioutil.WriteFile(vendorPriPath, []byte(content), 0666)
}

func (i *InstallCommand) loadSpmJson(spmJsonPath string, content *SpmJsonContent) error{
	var spmJsonConent interface{} = content
	err := util.LoadJsonFile(spmJsonPath, &spmJsonConent)
	if err!=nil {
		return err
	}
	return nil
}

//updateSpmJson 更新spm.json文件，若更新的包存在于spm.json文件中，则更新失败
func (i *InstallCommand) updateSpmJson(spmJsonPath string, content *SpmJsonContent) error{
	//检查原依赖文件中是否存在即将安装的包，存在时报错
	for _, dependency := range content.Dependencies {
		dependencyPkgInfo := strings.Split(dependency, "@")
		if dependencyPkgInfo==nil || len(dependencyPkgInfo)==0 {
			continue
		}
		if i.packageName==dependencyPkgInfo[0] {
			return errors.New("the same package already exists")
		}
	}
	content.Dependencies = append(content.Dependencies, i.packageName + "@" + i.version)
	return util.WriteStruct(spmJsonPath, content)
}

func (i *InstallCommand) checkDependency(content SpmJsonContent) error{
	installPackage := i.packageName + "@" + i.version
	for _, dependency := range content.Dependencies {
		if dependency==installPackage {
			return errors.New("the package installed")
		}
	}
	return nil
}

func NewInstallCommand() *InstallCommand{
	return &InstallCommand{}
}




