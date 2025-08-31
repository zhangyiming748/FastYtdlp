/*
在这里实现一个函数
每次main函数中下载成功一个文件
就把这次下载的url写入一个文本文件
在每次下载文件之前
提供一个函数用来检测文件是否被成功下载过文件
*/

package history

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"FastYTDLP/util"
)

const historyFile = "history.link"
const(
	ROOT = "/data"
	HISTORYFILE = "history.link"
)
// IsURLDownloaded 检查URL是否已经下载过
func IsURLDownloaded(url string) (bool) {
	history:=filepath.Join(ROOT, HISTORYFILE)
	links:=util.ReadByLine(history)
	for _, link := range links {
		/*
		如果link和url都包含"viewkey"那么进入比较
		*/
		if strings.Contains(link, "viewkey") && strings.Contains(url, "viewkey") {
			//根据等号拆分
			linkParts := strings.Split(link, "=")[1]
			if strings.Contains(linkParts,"#"){
				linkParts=strings.Split(linkParts,"#")[0]
			}
			urlParts := strings.Split(url, "=")[1]
			if strings.Contains(urlParts,"#"){
				urlParts=strings.Split(urlParts,"#")[0]
			}
			if linkParts == urlParts {
				return true
			}
		} else if link == url {
			return true
		}
	}

	
}

// RecordDownloadedURL 记录已下载的URL
func RecordDownloadedURL(url string) error {
	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(url + "\n")
	return err
}

// extractViewKey 从URL中提取viewkey参数
func extractViewKey(url string) string {
	// 查找viewkey参数
	start := strings.Index(url, "viewkey=")
	if start == -1 {
		return ""
	}
	start += len("viewkey=")
	
	// 查找参数结束位置（可能是&或者字符串结尾）
	end := strings.Index(url[start:], "&")
	if end == -1 {
		return url[start:]
	}
	
	return url[start : start+end]
}