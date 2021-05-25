package main

import (
	"auto_commit/config"
	"auto_commit/git"
	"github.com/kainhuck/gokit/sync"
)

func main() {
	var wg sync.WaitGroupWrapper
	for _, cfg := range config.Cfgs {
		wg.Wrap(func() {
			g := git.NewGit(cfg)
			g.Run()
		})
	}
	wg.Wait()
}
