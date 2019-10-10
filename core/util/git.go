package util

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func NewGit() *Git {
	return &Git{}
}

type Git struct {
}

func (g *Git) Test() error {
	_, err := exec.Command("git", "version").Output()
	if err != nil {
		return err
	}
	return nil
}

func (g *Git) CloneRepository(url string, destdir string) error {
	_, err := exec.Command("git", "clone", "--recursive", url, destdir).Output()
	if err != nil {
		return err
	}
	return nil
}

func (g *Git) CheckoutRevision(revision string) error {
	//log.Print("git checkout ", revision)
	_, err := exec.Command("git", "checkout", revision).Output()
	if err != nil {
		return err
	}
	return nil
}

func (g *Git) CreateTag(name string) error {
	_, err := exec.Command("git", "tag", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (g *Git) RepositoryURL() (string, error) {
	out, err := exec.Command("git", "config", "remote.origin.url").Output()
	if err != nil {
		return "", errors.New("we could not get the repository remote origin URL")
	}
	return strings.TrimSpace(string(out)), err
}

func (g *Git) LastCommitRevision() (string, error) {
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	return strings.TrimSpace(string(out)), err
}

func (g *Git) LastCommitAuthorName() (string, error) {
	args := []string{"log", "-1", "--format=%an"}
	out, err := exec.Command("git", args...).Output()
	return strings.TrimSpace(string(out)), err
}

func (g *Git) LastCommitEmail() (string, error) {
	args := []string{"log", "-1", "--format=%ae"}
	out, err := exec.Command("git", args...).Output()
	return strings.TrimSpace(string(out)), err
}


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

//IsGitRepository 判断指定的路径是否是一个git仓库
func IsGitRepository(dir string) bool {
	return IsExists(filepath.Join(dir, ".git"))
}
