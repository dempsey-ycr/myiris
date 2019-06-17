package main

import (
	"myiris/config"
	"myiris/service/web"
)

func main() {
	// 读取自定义配置
	config.Initialize()

	web.Iris()
}
