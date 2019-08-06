package util

import (
	"os"
	"os/exec"
	"path"
)

//GitClone 将git仓库源码clone到指定的路径下
func GitClone(url, p string) error{
	cmd := exec.Command("git", "clone", "--progress", url, p)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err!=nil {
		return err
	}
	return cmd.Wait()
}

//RemoveDotGit 删除指定路径下的.git目录
func RemoveDotGit(p string) error{
	dotGitPath := path.Join(p, ".git")
	if !IsExists(dotGitPath) {
		return nil
	}
	return os.RemoveAll(dotGitPath)
}
