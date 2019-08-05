package conf


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
}

type Repository struct {
	Url string	`json:"url"`
}

func NewSpmJson() *SpmJson{
	return &SpmJson{
		Version: "0.0.1",
		Author: &Author{},
		Dependencies: []string{},
		Repository: &Repository{},
	}
}