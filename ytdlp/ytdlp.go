package ytdlp

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/zhangyiming748/FastYtdlp/storage"
	"github.com/zhangyiming748/FastYtdlp/util"
)

func DownloadVideo(uri, proxy, location string) (title string) {
	if has, err := sameUrl(uri); err != nil {
		log.Fatalf("查询数据库失败:%v\n", err)
	} else if has {
		log.Printf("由于数据库中已存在相同链接%v\t跳过此次下载\n", uri)
		return uri
	}
	nameCmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, "--get-filename", uri)
	name := util.GetVideoName(nameCmd)
	name = filepath.Base(name)
	if has, err := sameName(uri); err != nil {
		log.Fatalf("查询数据库失败:%v\n", err)
	} else if has {
		log.Printf("由于数据库中已存在同名文件%v\t跳过此次下载\n", name)
		return uri
	}
	log.Printf("当前下载的文件标题:%s", name)
	downloadCmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, uri)
	util.ExecCommand4Ytdlp(downloadCmd)
	log.Printf("当前下载成功的文件标题:%s", name)
	one := new(storage.Video)
	one.Url = uri
	one.Name = name
	insertOne, err := one.InsertOne()
	if err != nil {
		log.Fatalf("插入%d条数据失败:%v\n", insertOne, err)
	} else {
		log.Printf("成功插入%d条数据\n", insertOne)
	}
	return name
}

func sameUrl(uri string) (bool, error) {
	one := new(storage.Video)
	one.Url = uri
	return one.FindByUrl()
}
func sameName(name string) (bool, error) {
	one := new(storage.Video)
	one.Name = name
	return one.FindByName()
}
