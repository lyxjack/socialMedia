package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xpzouying/xiaohongshu-mcp/configs"
)

func main() {
	var (
		headless bool
		binPath  string // 浏览器二进制文件路径
		port     string
		apiToken string // API 认证 Token
	)
	flag.BoolVar(&headless, "headless", true, "是否无头模式")
	flag.StringVar(&binPath, "bin", "", "浏览器二进制文件路径")
	flag.StringVar(&port, "port", "127.0.0.1:18060", "监听地址和端口")
	flag.StringVar(&apiToken, "token", "", "API 认证 Bearer Token")
	flag.Parse()

	if len(binPath) == 0 {
		binPath = os.Getenv("ROD_BROWSER_BIN")
	}

	if apiToken == "" {
		apiToken = os.Getenv("XHS_API_TOKEN")
	}
	if apiToken == "" {
		logrus.Warn("未配置 API Token，HTTP API 接口无认证保护。建议通过 -token 或 XHS_API_TOKEN 设置")
	}

	configs.InitHeadless(headless)
	configs.SetBinPath(binPath)

	// 初始化服务
	xiaohongshuService := NewXiaohongshuService()

	// 创建并启动应用服务器
	appServer := NewAppServer(xiaohongshuService)
	appServer.apiToken = apiToken
	if err := appServer.Start(port); err != nil {
		logrus.Fatalf("failed to run server: %v", err)
	}
}
