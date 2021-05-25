package utils

import (
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	t.Log(path.Join(HomeDir(), ".config/auto_commit/config.json"))
}
