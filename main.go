package main

import (
	"FastYTDLP/history"
	"FastYTDLP/util"
	"FastYTDLP/ytdlp"
	"log"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		proxy = "192.168.5.2:8889"
	}
	location := "/videos"
	lines := util.ReadByLine("/data/post.link")
	for i, line := range lines {
		log.Printf("正在处理第%d个链接:%s", i+1, line)
		if strings.Contains(line, "#") {
			prefix := strings.Split(line, "#")[0]
			suffix := strings.Split(line, "#")[1]
			local:=filepath.Join(location, suffix)
			os.MkdirAll(local, os.ModePerm)
			download(prefix, proxy, local)
		} else {
			download(line, proxy, location)
		}
	}
}

func download(line, proxy, location string) {
	if has, err := history.IsURLDownloaded(line); err != nil {
		log.Fatalf("发生需要立刻终止的严重错误:%v\n", err)
	} else if has {
		log.Printf("该链接已经下载过了,跳过:%s", line)
		return
	} else {
		ytdlp.DownloadVideo(line, proxy, location)
		history.RecordDownloadedURL(line)
	}
}
