package main

import "fmt"

/*
//接口的定义
type interface_name interface{
    function_name( [params] ) [return_values]
    ...
}
*/

// ImageDownloader 图片加载接口
type ImageDownloader interface {
	// FetchImage 获取图片，需要传入图片地址，方法返回图片数据
	FetchImage(url string) string
}

/*
	func (struct_variable struct_name) function_name([params]) [return_values] {
	   // 方法实现
	}
*/
type fileCache struct {
}

func (f *fileCache) FetchImage(url string) string {
	return "从本地缓存中获取图片：" + url
}

type netFetch struct {
}

func (n *netFetch) FetchImage(url string) string {
	return "从网络下载图片：" + url
}

func main() {
	imageLoader := new(fileCache)
	data := imageLoader.FetchImage("cache https://www.example.com/a.png")
	if data == "" {
		imageNetLoader := new(netFetch)
		data = imageNetLoader.FetchImage("cache https://www.example.com/a.png")
	}
	fmt.Println(data)
}
