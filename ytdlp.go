package FastYtdlp

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/FastYtdlp/storage"
	"github.com/zhangyiming748/FastYtdlp/util"
	"github.com/zhangyiming748/FastYtdlp/ytdlp"
)

func Download(root string, yc YtdlpConfig) {
	storage.SetMysql(yc.User, yc.Password, yc.Host, yc.Port)
	storage.GetMysql().Sync2(storage.Pornhub{})
	os.MkdirAll(root, os.ModePerm)
	post := filepath.Join(root, "post.link")
	log.Printf("根据%v文件开始下载\n", post)
	lines := util.ReadByLine(post)

	for i, line := range lines {
		log.Printf("正在处理第%d个链接:%s", i+1, line)

		if strings.Contains(line, "pornhub") {
			var key string
			if strings.Contains(line, "#") {
				prefix := strings.Split(line, "#")[0]
				key = strings.Split(prefix, "=")[1]
				suffix := strings.Split(line, "#")[1]
				local := filepath.Join(root, suffix)
				name := ytdlp.DownloadVideo(prefix, yc.Proxy, local)
				one := new(storage.Pornhub)
				one.Key = key
				if has, _ := one.FindByKey(); has {
					log.Printf("由于数据库中已存在%n\t跳过此次下载\n", one.Name)
					continue
				}
				one.Name = name
				one.From = "pornhub"
				one.InsertOne()
			} else {
				name := ytdlp.DownloadVideo(line, yc.Proxy, root)
				key = strings.Split(line, "=")[1]
				one := new(storage.Pornhub)
				one.Key = key
				if has, _ := one.FindByKey(); has {
					log.Printf("由于数据库中已存在%n\t跳过此次下载\n", one.Name)
					continue
				}
				one.Name = name
				one.From = "pornhub"
				one.InsertOne()
			}
		} else {
			//其他来源的网站
		}
	}
}
