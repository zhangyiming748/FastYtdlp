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
	storage.GetMysql().Sync2(storage.Video{})
	os.MkdirAll(root, os.ModePerm)
	post := filepath.Join(root, "post.link")
	log.Printf("根据%v文件开始下载\n", post)
	lines := util.ReadByLine(post)

	for i, line := range lines {
		log.Printf("正在处理第%d个链接:%s", i+1, line)

		if strings.Contains(line, "pornhub") {
			var key string
			if strings.Contains(line, "#") {
				uri := strings.Split(line, "#")[0]
				key = strings.Split(uri, "=")[1]
				subFolder := strings.Split(line, "#")[1]
				local := filepath.Join(root, subFolder)
				name := ytdlp.DownloadVideo(uri, yc.Proxy, local)
				one := new(storage.Video)
				one.Keyword = key
				if has, err := one.FindByKeyword(); err != nil {
					log.Fatalf("查询数据库失败:%v\n", err)
				} else if has {
					log.Printf("由于数据库中已存在%v\t跳过此次下载\n", one.Name)
					continue
				}
				one.Name = name
				one.From = "pornhub"
				insertOne, err := one.InsertOne()
				if err != nil {
					log.Printf("插入%d条数据失败:%v\n", insertOne, err)
				} else {
					log.Printf("成功插入%d条数据\n", insertOne)
				}
			} else {
				name := ytdlp.DownloadVideo(line, yc.Proxy, root)
				key = strings.Split(line, "=")[1]
				one := new(storage.Video)
				one.Keyword = key
				if has, err := one.FindByKeyword(); err != nil {
					log.Fatalf("查询数据库失败:%v\n", err)
				} else if has {
					log.Printf("由于数据库中已存在%v\t跳过此次下载\n", one.Name)
					continue
				}
				one.Name = name
				one.From = "pornhub"
				insertOne, err := one.InsertOne()
				if err != nil {
					log.Printf("插入%d条数据失败:%v\n", insertOne, err)
				} else {
					log.Printf("成功插入%d条数据\n", insertOne)
				}
			}
		} else {
			//其他来源的网站
		}
	}
}
