package storage

import (
	"log"
	"testing"
)

func TestGetMysql(t *testing.T) {
	SetMysql("root", "163453", "192.168.5.2", "3306")
}

func TestHasKey(t *testing.T) {
	SetMysql("root", "163453", "192.168.5.2", "3306")
	one := new(Pornhub)
	one.Key = "ph637a604fe4704"
	if has, err := one.FindByKey(); has {
		log.Printf("由于数据库中已存在%v\t跳过此次下载\n%v", one.Name, err)
	} else {
		log.Printf("数据库中不存在%v\t开始下载\n%v", one.Key, err)
	}
}
