package FastYtdlp

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/FastYtdlp/storage"
	"github.com/zhangyiming748/FastYtdlp/util"
)

func Download(root string, yc YtdlpConfig) {
	storage.SetMysql(yc.User, yc.Password, yc.Host, yc.Port)
	err := storage.GetMysql().Sync2(storage.Video{})
	if err != nil {
		log.Fatalf("创建数据库同步表结构连接失败:%v\n", err)
	}
	post := filepath.Join(root, "post.link")
	log.Printf("根据%v文件开始下载\n", post)
	lines := util.ReadByLine(post)
	for i, line := range lines {
		log.Printf("正在处理第%d个链接:%s", i+1, line)
		var name string
		if strings.Contains(line, "#") {
			uri := strings.Split(line, "#")[0]
			hashTag := strings.Split(line, "#")[1]
			local := filepath.Join(root, hashTag)
			name = DownloadHelper(uri, yc.Proxy, local)
		} else {
			name = DownloadHelper(line, yc.Proxy, root)
		}
		log.Printf("下载%v\n流程结束", name)
	}
}
func DownloadHelper(uri, proxy, location string) (title string) {
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
