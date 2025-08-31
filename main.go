package main

import (
	"FastYTDLP/history"
	"FastYTDLP/util"
	"FastYTDLP/ytdlp"
	"log"
	"os"
	"strings"
	"path/filepath"
	"io"
)
func init(){
	/*
	这里设置log包含 短时间 短文件名
	并且设置log同时输出到控制台和ytdlp.log文件中
	*/
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	
	// 创建或打开日志文件
	file, err := os.OpenFile("ytdlp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	
	// 设置log同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
}
func main() {
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		proxy = "192.168.5.2:8889"
	}
	location := "/data/videos"
	os.MkdirAll(location, os.ModePerm)
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
	if has:= history.IsURLDownloaded(line); has {
		log.Printf("该链接已经下载过了,跳过:%s", line)
		return
	} else {
		ytdlp.DownloadVideo(line, proxy, location)
		history.RecordDownloadedURL(line)
	}
}