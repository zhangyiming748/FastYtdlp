package storage

import "testing"

func TestGetMysql(t *testing.T) {
	SetMysql("root", "163453", "192.168.5.2", "3306")
}
