package conf

import (
	"spm/core/util"
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
	var data interface{} = s
	return util.LoadJsonFile(filePath, &data)
}

func NewSpmJson() *SpmJson{
	return &SpmJson{
		Version: "0.0.1",
		Author: &Author{},
		Dependencies: []string{},
		Repository: &Repository{},
	}
}