package config

import (
	"auto_commit/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	Cfgs        = make([]*Config, 0)
	TplFileName = "config.json.tpl" // 配置文件名
	ConfPath    string              // 配置文件路径
	FullPath    string              // 配置文件完整路径
)

func init() {
	ConfPath = path.Join(utils.HomeDir(), ".config/auto_commit")
	FullPath = path.Join(ConfPath, TplFileName)
	//检测模板配置文件是否存在
	if !utils.CheckFileExist(FullPath) {
		MkTplConfig()
	}
	LoadCfgs()
}

type Config struct {
	ProjectRootPath string `json:"project_root_path"`
	CommitMsg       string `json:"commit_msg"`     // todo
	CheckInterval   int    `json:"check_interval"` // 单位毫秒
	Username        string `json:"username"`
	Watching        bool   `json:"watching"` // 是否监视该项目
}

func (c *Config) Load(fileName string) {
	bts, err := os.ReadFile(path.Join(ConfPath, fileName))
	if err != nil {
		log.Fatalf("LoadConfig error: %v", err)
	}

	err = json.Unmarshal(bts, &c)
	if err != nil {
		log.Fatalf("LoadConfig error: %v", err)
	}
}

func MkTplConfig() {
	_ = os.MkdirAll(ConfPath, os.ModePerm)
	f, err := os.Create(FullPath)
	if err != nil {
		log.Fatalf("Create config file failed: %v", err)
	}
	defer f.Close()

	defaultCfg := Config{
		ProjectRootPath: "Your/Project/Path",
		CommitMsg:       "%d commit by auto_commit",
		CheckInterval:   10000,
		Username:        "your_name",
		Watching:        false,
	}
	bts, _ := json.MarshalIndent(defaultCfg, "", "	")
	f.Write(bts)
}

// 加载配置文件
func LoadCfgs() {
	// 列出配置目录下的所有json文件
	fs, _ := ioutil.ReadDir(ConfPath)
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			cfg := new(Config)
			cfg.Load(f.Name())
			Cfgs = append(Cfgs, cfg)
		}
	}

	return
}
