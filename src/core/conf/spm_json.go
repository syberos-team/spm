package conf

import (
	"core/util"
	"encoding/json"
	"io/ioutil"
	"os"
)

type SpmJson struct {
	Name string 	`json:"name"`
	Description string	`json:"description"`
	Version string	`json:"version"`
	Author *Author	`json:"author"`
	Dependencies []string	`json:"dependencies"`
	Repository *Repository	`json:"repository"`
	PriFilename string `json:"priFilename"`
}

type Author struct {
	Name string		`json:"name"`
	Email string	`json:"email"`
	Description string `json:"description"`
}

type Repository struct {
	Url string	`json:"url"`
}

//Load 加载spm.json文件信息
func (s *SpmJson) Load(filePath string) error{
	file, err := os.Open(filePath)
	if err!=nil {
		return err
	}
	defer util.CloseQuietly(file)
	data, err := ioutil.ReadAll(file)
	if err!=nil {
		return err
	}
	return json.Unmarshal(data, s)
}

func NewSpmJson() *SpmJson{
	return &SpmJson{
		Version: "0.0.1",
		Author: &Author{},
		Dependencies: []string{},
		Repository: &Repository{},
	}
}