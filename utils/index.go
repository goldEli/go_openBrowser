package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// OpenBrowser 使用 chromedp 模拟浏览器打开指定的URL
func OpenBrowser(url string) error {

	// opts := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.NoDefaultBrowserCheck,                   //不检查默认浏览器
	// 	chromedp.Flag("headless", false),                 //开启图像界面,重点是开启这个
	// 	chromedp.Flag("ignore-certificate-errors", true), //忽略错误
	// 	chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
	// 	chromedp.NoFirstRun,                              //设置网站不是首次运行
	// 	chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
	// )
	// allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// //	defer cancel()
	// ctx, _ := chromedp.NewContext(
	// 	allocCtx,
	// 	chromedp.WithLogf(log.Printf),
	// )

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// navigate to a page, wait for an element, click
	// var example string
	println("打开浏览器:", url)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		//chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		//chromedp.Value(`#example-After textarea`, &example),
	)

	if err != nil {
		fmt.Printf("发生错误：%v\n", err)
	}
	println("浏览器已打开")
	//睡眠20秒后退出
	time.Sleep(10 * time.Second)
	println("关闭浏览器:", url)
	// log.Printf("Go's time.After example:\n%s", example)
	return nil
}
