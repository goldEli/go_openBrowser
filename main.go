package main

import (
	"fmt"
	"net/http"
	"openBrowserTool/utils"
)

func handleDownload(w http.ResponseWriter, r *http.Request) {
	// 获取名为 "url" 的查询参数
	urlParam := r.URL.Query().Get("url")

	if urlParam == "" {
		http.Error(w, "参数 'url' 不能为空", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "url:", urlParam)
	print("url:", urlParam)
	print("start===>")
	utils.OpenBrowser(urlParam)
}

func main() {
	// 设置路由
	http.HandleFunc("/download", handleDownload)

	// 启动HTTP服务，监听在8080端口
	fmt.Println("服务启动，监听在 http://localhost:8080/download")
	http.ListenAndServe(":8080", nil)
}
