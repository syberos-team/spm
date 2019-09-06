package conf

import (
	"os"
	"os/user"
	"path"
	"spm/core/log"
	"spm/core/util"
)

const (
	ConfigDir string = ".spm"
	ConfigFilename string = "spm.conf"
)

var Config configInfo

func InitConfig(){
	Config = configInfo{
		Url: ServerUrl,
		Log: logInfo{
			Level:log.LogInfoString,
		},
	}
	Config.load()
	//设置日志级别
	log.SetLogLevelString(Config.Log.Level)
}

//日志配置信息
type logInfo struct {
	//日志级别
	Level string `json:"level"`
}

//配置信息
type configInfo struct {
	//服务器路径
	Url string `json:"url"`
	//日志
	Log logInfo `json:"log"`

	configDir string		//配置文件所在目录
	configFilePath string	//配置文件所在路径
}

//GetConfigDir 获取配置文件所在目录
func (c *configInfo) GetConfigDir() string{
	return c.configDir
}

//load 尝试从用户目录下加载.spm/spm.conf文件，若文件不存在，将会尝试创建，并使用默认配置
func (c *configInfo) load(){
	configFilePath, err := c.spmPath(ConfigFilename)
	c.configFilePath = configFilePath
	if err!=nil {
		log.Warning(err.Error())
		return
	}
	err = c.createSpmConf()
	if err!=nil {
		log.Error(err.Error())
		panic("failed to create spm.conf file")
	}
	err = c.readConfig(configFilePath)
	if err!=nil {
		log.Error(err.Error())
		panic("failed to read spm.conf file")
	}
}


//读取配置文件信息
func (c *configInfo) readConfig(configPath string) error {
	var data interface{} = c
	return util.LoadJsonFile(configPath, &data)
}

//spmPath 获取.spm目录下的子路径
func (c *configInfo) spmPath(sub ...string) (string, error){
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	c.configDir = path.Join(usr.HomeDir, ConfigDir)
	if sub==nil {
		return c.configDir, nil
	} else {
		dirs := []string{c.configDir}
		dirs = append(dirs, sub...)
		return 	path.Join(dirs...), nil
	}
}

//创建spm.conf文件
func (c *configInfo) createSpmConf() error{
	if util.IsExists(c.configFilePath) {
		return nil
	}
	log.Debug("no spm.conf file found, ready to create:", c.configFilePath)
	if err := os.MkdirAll(path.Dir(c.configFilePath), os.ModePerm); err!=nil {
		return err
	}
	return util.WriteStruct(c.configFilePath, c)
}





