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
	"strings"
)

const historyFile = "download_history.txt"

// IsURLDownloaded 检查URL是否已经下载过
func IsURLDownloaded(url string) (bool, error) {
	file, err := os.Open(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			//创建这个文件
			os.Create(historyFile)
		}
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == url {
			return true, nil
		}
	}

	return false, scanner.Err()
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
