package git

import (
	"auto_commit/config"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Git struct {
	times int // 提交次数
	Cfg   *config.Config
}

func NewGit(cfg *config.Config) *Git {
	return &Git{
		Cfg: cfg,
	}
}

func (g *Git) SetName() error{
	cmd := exec.Command("git", "config", "--local", "user.name", g.Cfg.Username)
	cmd.Dir = g.Cfg.ProjectRootPath
	return cmd.Run()
}

// 判断项目路径是否存在且为git目录
func (g *Git) Status() error {
	cmd := exec.Command("git", "status")
	cmd.Dir = g.Cfg.ProjectRootPath
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// git add .
func (g *Git) Add() error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = g.Cfg.ProjectRootPath
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// git commit -m "..."
func (g *Git) Commit() error {
	g.times++
	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf(g.Cfg.CommitMsg, g.times))
	cmd.Dir = g.Cfg.ProjectRootPath
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		g.times--
	}
	return err
}

// git push
func (g *Git) Push() error {
	cmd := exec.Command("git", "push")
	cmd.Dir = g.Cfg.ProjectRootPath
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (g *Git) Run() {
	tick := time.NewTicker(time.Duration(g.Cfg.CheckInterval) * time.Millisecond)
	for {
		select {
		case <-tick.C:
			g.SetName()
			g.Status()
			g.Add()
			g.Commit()
			g.Push()
			g.Status()
		}
	}
}
