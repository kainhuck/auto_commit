package git

import (
	"auto_commit/config"
	"testing"
)

var g = NewGit(&config.Config{
	ProjectRootPath: "/Users/kainhuck/go/src/fundgo",
	CommitMsg:       "commit by auto_commit %d",
	CheckInterval:   3000,
})

func TestGit_Status(t *testing.T) {
	g.Status()
}

func TestGit_Add(t *testing.T) {
	g.Add()
}

func TestGit_Commit(t *testing.T) {
	g.Commit()
}

func TestGit_Push(t *testing.T) {
	g.Push()
}

func TestGit_GetOldName(t *testing.T) {
	g.GetOldName()
	t.Log(g.oldName)
}
