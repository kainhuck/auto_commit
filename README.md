# auto_commit
用于自动提交代码到git

# 如何使用
首次运行该文件时，会创建`~/.config/auto_commit/config.json.tpl`模版文件

你需要做的就是通过该文件创建一个项目配置文件，注意必须在`~/.config/auto_commit`目录下，必须是`.json`结尾的文件名。

模版文件介绍：
```json
{
  "project_root_path": "Your/Project/Path",
  "commit_msg": "%d commit by auto_commit",
  "check_interval": 3000,
  "username": "your_name",
  "watching": false
}
```
- project_root_path
  项目根目录
- commit_msg
  提交信息，%d 会传入当前提交的次数
- check_interval
  检查提交的间隔，单位毫秒
- username
  提交用户名
- watching
  是否监视该项目，如果为false则不监控