package commands

import (
	"core/util"
	"errors"
	"flag"
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
	SpmJson = "spm.json"
)

type InstallCommand struct {
	Command
	packageName string
	version string
}

func (i *InstallCommand) Description() string {
	return "Installs a new package"
}

func (i *InstallCommand) Run() error {
	spmJsonContent := &SpmJsonContent{}

	//检查spm.json文件中是否已存在待安装的包
	spmJsonFilePath := path.Join(pwd, spmJsonFilename)
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

func (i *InstallCommand) RegisterFlags(flags *flag.FlagSet) {
	arg := flags.Arg(0)
	if arg=="" {
		return
	}
	packageInfo := strings.Split(arg, "@")
	if packageInfo==nil || len(packageInfo)==0{
		return
	}
	packageInfoLen := len(packageInfo)
	if packageInfoLen==1{
		i.packageName = packageInfo[0]
	}else if packageInfoLen==2 {
		i.packageName = packageInfo[0]
		i.version = packageInfo[1]
	}
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

func (i *InstallCommand) updateVendor(vendorPriPath string) error{
	file, err := os.Open(vendorPriPath)
	if err!=nil {
		return err
	}
	defer util.CloseQuietly(file)


	return nil
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




