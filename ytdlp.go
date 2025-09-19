package FastYtdlp

import (
	"github.com/zhangyiming748/FastYtdlp/util"
	"github.com/zhangyiming748/FastYtdlp/ytdlp"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Ytdlp(root, proxy string) {
	os.MkdirAll(root, os.ModePerm)
	post:=filepath.Join(root, "post.link")
	lines := util.ReadByLine(post)
	for i, line := range lines {
		log.Printf("正在处理第%d个链接:%s", i+1, line)
		if strings.Contains(line, "#") {
			prefix := strings.Split(line, "#")[0]
			suffix := strings.Split(line, "#")[1]
			local := filepath.Join(root, suffix)
			ytdlp.DownloadVideo(prefix, proxy, local)
		} else {
			ytdlp.DownloadVideo(line, proxy, root)
		}
	}
}

