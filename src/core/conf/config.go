package conf

import (
	"core/util"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const (
	ConfigDir string = ".spm"
	ConfigFilename string = "spm.conf"
	ServerUrl string = "http://www.baidu.com"
)

var Config configInfo

type configInfo struct {
	Url string `json:"url"`
	configPath string
}

func init(){
	Config = configInfo{}
	Config.load()
}

func (c *configInfo) load(){
	configPath, err := c.spmPath(ConfigFilename)
	c.configPath = configPath
	if err!=nil {
		panic(err.Error())
	}
	if !util.IsExists(configPath) {
		c.Url = ServerUrl
		return
	}
	err = c.readConfig(configPath)
	if err!=nil {
		panic("failed to read spm.conf file")
	}
}


func (c *configInfo) readConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err !=nil {
		return err
	}
	defer util.CloseQuietly(file)
	content, err := ioutil.ReadAll(file)
	if err!=nil{
		return err
	}
	err = json.Unmarshal(content, c)
	if err!=nil {
		return err
	}
	return nil
}

//spmPath 获取.spm目录下的子路径
func (c *configInfo) spmPath(sub ...string) (string, error){
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	if sub==nil {
		return path.Join(usr.HomeDir, ConfigDir), nil
	} else {
		dirs := []string{usr.HomeDir, ConfigDir}
		dirs = append(dirs, sub...)
		return 	path.Join(dirs...), nil
	}
}

func (c *configInfo) mkSpmConf() error{
	if util.IsExists(c.configPath) {
		return nil
	}
	if err := os.MkdirAll(path.Dir(c.configPath), os.ModePerm); err!=nil {
		return err
	}
	c.Url, _ = <-util.Prompt("server interface address:", "")

	file, err := os.Create(c.configPath)
	if err!=nil {
		return err
	}
	defer util.CloseQuietly(file)
	if data, err := util.ToPrettyJSON(*c); err==nil {
		_, err := file.Write(data)
		return err
	}else{
		return err
	}
}





