package conf

import "io"

//TODO 待实现操作spm执行文件的创建和覆盖

const (
	SpmFilename = "spm"
	TempSpmFilename = "spm.temp"
)

type SpmFile struct {
	//父级目录路径
	parentPath string
}

func (s *SpmFile) Create(w *io.Writer){

}

func (s *SpmFile) Replace() {

}

func NewSpmFile() SpmFile{
	s := SpmFile{}

	return s
}


