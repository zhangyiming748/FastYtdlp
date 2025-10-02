package FastYtdlp

import (
	"testing"
)

func TestYtdlp(t *testing.T) {
	root := "C:\\Users\\zen\\Github\\FastYtdlp"
	var yc YtdlpConfig
	yc.Host = "192.168.5.2"
	yc.Port = "3306"
	yc.User = "root"
	yc.Password = "163453"
	yc.Proxy = "192.168.5.2:8889"
	Download(root, yc)
}
