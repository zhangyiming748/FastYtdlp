package FastYtdlp

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/FastYtdlp/sqlite"
	"github.com/zhangyiming748/FastYtdlp/util"
)

func init() {
	sqlite.SetSqlite()
	//在这里写一个同步表结构的代码
	new(sqlite.YtdlpHistory).Sync()
}
func Download(root, proxy string) {
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
			name = DownloadHelper(uri, proxy, local)
		} else {
			name = DownloadHelper(line, proxy, root)
		}
		log.Printf("下载%v\n流程结束", name)
	}
}
func DownloadHelper(uri, proxy, location string) (title string) {
	if has, _ := sameUrl(uri); has {
		log.Printf("由于数据库中已存在相同链接%v\t跳过此次下载\n", uri)
		return uri
	}
	nameCmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, "--get-filename", uri)
	name := util.GetVideoName(nameCmd)
	name = filepath.Base(name)
	if has, _ := sameName(name); has {
		log.Printf("由于数据库中已存在同名文件%v\t跳过此次下载\n", name)
		return name
	}
	log.Printf("当前下载的文件标题:%s", name)
	downloadCmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, uri)
	util.ExecCommand4Ytdlp(downloadCmd)
	log.Printf("当前下载成功的文件标题:%s", name)
	one := new(sqlite.YtdlpHistory)
	one.Url = uri
	one.Name = name
	insertOne, err := one.InsertOne()
	if err != nil {
		log.Fatalf("插入%d条数据失败:%v\n", insertOne, err)
	}
	log.Printf("成功插入%d条数据\n", insertOne)
	return name
}

func sameUrl(uri string) (bool, error) {
	one := new(sqlite.YtdlpHistory)
	one.Url = uri
	return one.ExistsByUrl()
}
func sameName(name string) (bool, error) {
	one := new(sqlite.YtdlpHistory)
	one.Name = name
	return one.ExistsByName()
}
