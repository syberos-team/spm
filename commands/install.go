package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"spm/core/log"
	"spm/core/util"
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
	packageName string		//包名，通过参数获得，若参数中没有待安装的包，将使用当前目录下的spm.json中配置的依赖安装
	version string		//版本号，通过参数获得，若参数中没有是，通过接口获取

	priFilename string 		//pri文件名，通过接口获取

	installAll bool  //是否本地所有依赖，若参数中没有指定待安装的包时为true
}

func (i *InstallCommand) Description() string {
	return "Installs a new package"
}

func (i *InstallCommand) Run() error {
	log.Debug("install package:", i.packageName, i.version)

	// 当前目录下的spm.json路径
	spmJsonFilePath := path.Join(pwd, SpmJsonFilename)

	// 1 执行命令时未指定包，表示需要安装spm.json中的依赖
	if i.installAll {
		log.Debug("install all packages: ", spmJsonFilePath)
		// spm.json文件不存在，抛错
		if !util.IsExists(spmJsonFilePath) {
			log.Debug("no spm.json file found: ", spmJsonFilePath)
			return errors.New("no spm.json file found")
		}
		spmJsonContent, err := i.loadSpmJson(spmJsonFilePath)
		if err!=nil {
			return err
		}
		pkgs := i.findDependencies(spmJsonContent)
		// 1.1 对比vendor目录下的包与spm.json中的依赖包，不存在需要安装的包，结束安装
		if pkgs==nil || len(pkgs)==0 {
			return nil
		// 1.2 对比vendor目录下的包与spm.json中的依赖包，存在需要安装的包，开始安装包
		}else{
			for _, pkg := range pkgs {
				pkgName, pkgVersion := util.ParsePackageInfo(pkg)
				err := i.installPackage(spmJsonFilePath, pkgName, pkgVersion, spmJsonContent, false)
				if err!=nil {
					return err
				}
			}
		}
	// 2 执行命令时指定了包
	} else {
		log.Debug("install one package: ", i.packageName, i.version)
		// 2.1 当前目录下存在spm.json
		if util.IsExists(spmJsonFilePath) {
			spmJsonContent, err := i.loadSpmJson(spmJsonFilePath)
			if err!=nil {
				return err
			}
			// 2.1.1 spm.json中存在指定的包，抛错：包已安装
			if i.isPkgInSpmJson(spmJsonContent) {
				return errors.New("the package is installed")

			// 2.1.2 spm.json中不存在指定的包，开始安装包并更新spm.json
			}else{
				err := i.installPackage(spmJsonFilePath, i.packageName, i.version, spmJsonContent, true)
				if err!=nil {
					return err
				}
			}
		// 2.2 当前目录下不存在spm.json，开始安装包并创建spm.json
		} else {
			// 检查需要安装的包是否存在于vendor目录下
			if isExists, pkgSpmJsonPath := i.isDependencyExists(i.packageName); isExists {
				log.Debug("the package in vendor is exists: ", pkgSpmJsonPath)
				return errors.New("the package in vendor is exists")
			}

			err := i.installPackage(spmJsonFilePath, i.packageName, i.version, nil, true)
			if err!=nil {
				return err
			}
		}
	}
	return nil
}

func (i *InstallCommand) RegisterArgs(args ...string) {
	if len(args) > 0 {
		i.packageName, i.version = util.ParsePackageInfo(args[0])
	}else{
		i.installAll = true
	}
}

func (i *InstallCommand) ArgsDescription() []ArgsDescription{
	return []ArgsDescription{
		{
			Name:        "package",
			Description: "The package name can be attached with a version number, such as com.syber.test@1.0.0",
			Required:    false,
			IsArray:     false,
		},
	}
}

// 安装包，若spm.json不存在则会创建，参数spmJsonContent表示spm.json的内容，故当spm.json不存在时为空
func (i *InstallCommand) installPackage(spmJsonFilePath, pkgName, pkgVersion string, spmJsonContent *SpmJsonContent, needUpdateSpmJson bool) error{
	//通过查询待安装的包信息，infoData由接口返回
	infoData, err := PackageInfo(pkgName, pkgVersion)
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
	subDir := path.Join(vendorDirPath, i.pkgName2Path(infoData.Package.Name))
	err = i.downloadFromGit(repoUrl, subDir)

	//更新spm.json
	if needUpdateSpmJson {
		err = i.updateSpmJson(spmJsonFilePath, infoData.Package.Name, infoData.Version, spmJsonContent)
		if err!=nil {
			return err
		}
	}
	//更新vendor.pri
	err = i.updateVendor(infoData.Package.Name, path.Join(vendorDirPath, VendorPri))
	if err!=nil {
		return err
	}
	return nil
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

func (i *InstallCommand) createVendorPriContent(pkgName, vendorPriPath string) (string, error){
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
		path.Join(i.pkgName2Path(pkgName), i.priFilename) +
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
func (i *InstallCommand) updateVendor(pkgName, vendorPriPath string) error{
	content, err := i.createVendorPriContent(pkgName, vendorPriPath)
	if err!=nil {
		return err
	}
	return ioutil.WriteFile(vendorPriPath, []byte(content), 0666)
}

func (i *InstallCommand) loadSpmJson(spmJsonPath string) (*SpmJsonContent, error){
	var spmJsonContent = &SpmJsonContent{}
	err := util.LoadJsonFile(spmJsonPath, spmJsonContent)
	if err!=nil {
		return nil, err
	}
	return spmJsonContent, nil
}

//updateSpmJson 更新spm.json文件，若参数content为空，表示spm.json文件不存在，故此时是创建新的spm.json文件
func (i *InstallCommand) updateSpmJson(spmJsonPath, pkgName, pkgVersion string, content *SpmJsonContent) error{
	if content==nil {
		content = &SpmJsonContent{}
	}
	content.Dependencies = append(content.Dependencies, pkgName + "@" + pkgVersion)
	return util.WriteStruct(spmJsonPath, content)
}

// 检查安装包是否存在于spm.json文件中
func (i *InstallCommand) isPkgInSpmJson(content *SpmJsonContent) bool{
	hasVersion := i.version!=""
	installPackage := i.packageName + "@" + i.version

	if log.IsDebug() {
		log.Debug("check dependency, install:", installPackage, "dependencies[spm.json]:", fmt.Sprintf("%+v",content.Dependencies))
	}
	for _, dependency := range content.Dependencies {
		if hasVersion {
			return dependency==installPackage
		}
		return strings.HasPrefix(dependency, installPackage)
	}
	return false
}

// 对比vendor目录下和spm.json中的依赖，找到所有spm.json中未安装的包，使用@分割包名和版本号
func (i *InstallCommand) findDependencies(content *SpmJsonContent) []string{
	if content.Dependencies == nil || len(content.Dependencies)==0 {
		return nil
	}
	// 记录待安装的包
	var pkgs []string

	for _, dependency := range content.Dependencies {
		pkgName, _ := util.ParsePackageInfo(dependency)

		// vendor的指定包的路径下未找到spm.json表示没有安装该包
		if isExist, spmJsonPath := i.isDependencyExists(pkgName); !isExist {
			log.Debug("spm.json not exist: ", spmJsonPath)
			pkgs = append(pkgs, pkgName)
			continue
		}
	}
	return pkgs
}

// 依赖包是否已存在于vendor目录下
// 返回 是否存在、vendor中指定包目录下的spm.json路径
func (i *InstallCommand) isDependencyExists(pkgName string) (isExist bool, spmJsonPath string){
	spmJsonPath = filepath.Join(pwd, Vendor, i.pkgName2Path(pkgName), SpmJsonFilename)
	return util.IsExists(spmJsonPath), spmJsonPath
}

// 将包名转换成路径形式
func (i *InstallCommand) pkgName2Path(pkgName string) string{
	return strings.ReplaceAll(pkgName, ".", "/")
}

func NewInstallCommand() *InstallCommand{
	return &InstallCommand{}
}




