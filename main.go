package main

import (
	"flag"
	"fmt"
	"openBrowserTool/utils"
	"os"
)

func main() {
	// 定义命令行参数
	urlFlag := flag.String("url", "", "URL to open in the browser")

	// 解析命令行参数
	flag.Parse()

	// 获取解析后的参数值
	url := *urlFlag

	if url == "" {
		fmt.Println("请提供要打开的URL，使用 -url 参数")
		os.Exit(1)
	}

	println(url)

	utils.OpenBrowser(url)
	// 打开默认浏览器
	// cmd := exec.Command("open", url)
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("无法打开浏览器:", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("正在打开浏览器：%s\n", url)
}
