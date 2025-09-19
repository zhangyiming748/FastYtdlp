package FastYtdlp

import (
	"testing"
)

func TestYtdlp(t *testing.T) {
	root := "C:\\Users\\zen\\Github\\FastYtdlp"
	proxy := "192.168.5.2:8889"
	Ytdlp(root, proxy)
}
